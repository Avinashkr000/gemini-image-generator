package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"gemini-image-generator/config"
	"gemini-image-generator/models"
)

type GeminiImageRequest struct {
	Prompt string `json:"prompt"`
}

type GeminiImageResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text       string `json:"text,omitempty"`
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
		Prompt: req.Prompt,
		Status: "pending",
	}

	// Insert into database
	if err := config.DB.Create(&image).Error; err != nil {
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
			"temperature":      1,
			"topK":             40,
			"topP":             0.95,
			"maxOutputTokens":  8192,
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
		config.DB.Model(&image).Updates(map[string]interface{}{
			"status":     "failed",
			"updated_at": time.Now(),
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call Gemini API: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		config.DB.Model(&image).Updates(map[string]interface{}{
			"status":     "failed",
			"updated_at": time.Now(),
		})
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Gemini API error (status %d): %s", resp.StatusCode, string(body)),
		})
		return
	}

	var geminiResp GeminiImageResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		config.DB.Model(&image).Updates(map[string]interface{}{
			"status":     "failed",
			"updated_at": time.Now(),
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
		config.DB.Model(&image).Updates(map[string]interface{}{
			"status":     "failed",
			"updated_at": time.Now(),
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No image generated in response",
			"debug": string(body),
		})
		return
	}

	// Update image record with generated image
	config.DB.Model(&image).Updates(map[string]interface{}{
		"image_url":  imageData,
		"status":     "completed",
		"updated_at": time.Now(),
	})

	// Fetch updated image
	config.DB.First(&image, image.ID)

	c.JSON(http.StatusOK, gin.H{
		"id":       image.ID,
		"prompt":   image.Prompt,
		"imageUrl": image.ImageURL,
		"status":   image.Status,
	})
}

func GetImages(c *gin.Context) {
	var images []models.Image

	// Find all images, sorted by creation date (newest first)
	if err := config.DB.Order("created_at DESC").Find(&images).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch images"})
		return
	}

	c.JSON(http.StatusOK, images)
}

func GetImageByID(c *gin.Context) {
	id := c.Param("id")
	imageID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	var image models.Image
	if err := config.DB.First(&image, imageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.JSON(http.StatusOK, image)
}

func DeleteImage(c *gin.Context) {
	id := c.Param("id")
	imageID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	var image models.Image
	if err := config.DB.First(&image, imageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	if err := config.DB.Delete(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
