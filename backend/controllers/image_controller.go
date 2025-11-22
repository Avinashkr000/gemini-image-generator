package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"gemini-image-generator/config"
	"gemini-image-generator/models"
)

var imageCollection *mongo.Collection

func init() {
	if config.DB != nil {
		imageCollection = config.DB.Collection("images")
	}
}

type GeminiImageRequest struct {
	Prompt string `json:"prompt"`
}

type GeminiImageResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text        string `json:"text,omitempty"`
				InlineData struct {
					MimeType string `json:"mimeType"`
					Data     string `json:"data"`
				} `json:"inlineData,omitempty"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func GenerateImage(c *gin.Context) {
	var req models.GenerateImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Create image record
	image := models.Image{
		ID:        primitive.NewObjectID(),
		Prompt:    req.Prompt,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert into database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if imageCollection == nil {
		imageCollection = config.DB.Collection("images")
	}

	_, err := imageCollection.InsertOne(ctx, image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image record"})
		return
	}

	// Call Gemini API
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gemini API key not configured"})
		return
	}

	// Using Gemini 2.0 Flash experimental model for image generation
	geminiURL := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-exp:generateContent?key=%s", apiKey)

	// Prepare request body with proper format for image generation
	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": fmt.Sprintf("Generate an image: %s", req.Prompt),
					},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"temperature":     1,
			"topK":            40,
			"topP":            0.95,
			"maxOutputTokens": 8192,
			"responseMimeType": "image/jpeg",
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare request"})
		return
	}

	resp, err := http.Post(geminiURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		// Update status to failed
		imageCollection.UpdateOne(ctx, bson.M{"_id": image.ID}, bson.M{
			"$set": bson.M{"status": "failed", "updatedAt": time.Now()},
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call Gemini API: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		imageCollection.UpdateOne(ctx, bson.M{"_id": image.ID}, bson.M{
			"$set": bson.M{"status": "failed", "updatedAt": time.Now()},
		})
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Gemini API error (status %d): %s", resp.StatusCode, string(body)),
		})
		return
	}

	var geminiResp GeminiImageResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		imageCollection.UpdateOne(ctx, bson.M{"_id": image.ID}, bson.M{
			"$set": bson.M{"status": "failed", "updatedAt": time.Now()},
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Gemini response"})
		return
	}

	// Extract image data from response
	var imageData string
	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		part := geminiResp.Candidates[0].Content.Parts[0]
		if part.InlineData.Data != "" {
			// Image data is base64 encoded
			imageData = fmt.Sprintf("data:%s;base64,%s", part.InlineData.MimeType, part.InlineData.Data)
		}
	}

	if imageData == "" {
		imageCollection.UpdateOne(ctx, bson.M{"_id": image.ID}, bson.M{
			"$set": bson.M{"status": "failed", "updatedAt": time.Now()},
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No image generated in response",
			"debug": string(body),
		})
		return
	}

	// Update image record with generated image
	image.ImageURL = imageData
	image.Status = "completed"
	image.UpdatedAt = time.Now()

	_, err = imageCollection.UpdateOne(ctx, bson.M{"_id": image.ID}, bson.M{
		"$set": bson.M{
			"imageUrl":  imageData,
			"status":    "completed",
			"updatedAt": time.Now(),
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       image.ID.Hex(),
		"prompt":   image.Prompt,
		"imageUrl": image.ImageURL,
		"status":   image.Status,
	})
}

func GetImages(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if imageCollection == nil {
		imageCollection = config.DB.Collection("images")
	}

	// Find all images, sorted by creation date (newest first)
	cursor, err := imageCollection.Find(ctx, bson.M{}, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch images"})
		return
	}
	defer cursor.Close(ctx)

	var images []models.Image
	if err = cursor.All(ctx, &images); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode images"})
		return
	}

	if images == nil {
		images = []models.Image{}
	}

	c.JSON(http.StatusOK, images)
}

func GetImageByID(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if imageCollection == nil {
		imageCollection = config.DB.Collection("images")
	}

	var image models.Image
	err = imageCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&image)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch image"})
		return
	}

	c.JSON(http.StatusOK, image)
}

func DeleteImage(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if imageCollection == nil {
		imageCollection = config.DB.Collection("images")
	}

	result, err := imageCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
