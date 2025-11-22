package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Prompt      string             `json:"prompt" bson:"prompt"`
	ImageURL    string             `json:"imageUrl" bson:"imageUrl"`
	Status      string             `json:"status" bson:"status"` // pending, completed, failed
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type GenerateImageRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type GenerateImageResponse struct {
	ID       string `json:"id"`
	Prompt   string `json:"prompt"`
	ImageURL string `json:"imageUrl"`
	Status   string `json:"status"`
}
