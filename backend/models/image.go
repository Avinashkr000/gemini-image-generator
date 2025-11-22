package models

import (
	"time"
)

type Image struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Prompt    string    `json:"prompt" gorm:"type:text;not null"`
	ImageURL  string    `json:"imageUrl" gorm:"type:longtext"`
	Status    string    `json:"status" gorm:"type:varchar(50);default:'pending'"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type GenerateImageRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type GenerateImageResponse struct {
	ID       uint   `json:"id"`
	Prompt   string `json:"prompt"`
	ImageURL string `json:"imageUrl"`
	Status   string `json:"status"`
}
