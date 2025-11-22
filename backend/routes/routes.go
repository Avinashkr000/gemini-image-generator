package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gemini-image-generator/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API routes
	api := router.Group("/api")
	{
		// Image generation routes
		images := api.Group("/images")
		{
			images.POST("/generate", controllers.GenerateImage)
			images.GET("", controllers.GetImages)
			images.GET("/:id", controllers.GetImageByID)
			images.DELETE("/:id", controllers.DeleteImage)
		}
	}
}
