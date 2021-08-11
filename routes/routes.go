package routes

import (
	"leozz37/glucose-measure-api/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.GET("/glucose", handlers.GetGlucose)
	router.NoRoute(noRoute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"status": 404, "message": "Page not found"})
}
