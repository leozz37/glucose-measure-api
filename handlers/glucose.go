package handlers

import (
	"leozz37/glucose-measure-api/middlewares"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetGlucose(c *gin.Context) {
	fileName := os.Getenv("FILE_NAME")
	fileDownloadURL := os.Getenv("FILE_DOWNLOAD_URL")
	err := middlewares.DownloadFile(fileName, fileDownloadURL)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": "Can't download file"})
	}

	middlewares.UnzipFile(fileName)
	glucoseLevel, err := middlewares.GetLastMeasureFromCSV(fileName)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": "Can't get glucose level from CSV"})
	}

	c.JSON(http.StatusOK, gin.H{"glucoseLevel": glucoseLevel, "status": 200})
}
