package handlers

import (
	"leozz37/glucose-measure-api/middlewares"
	"leozz37/glucose-measure-api/models"
	"leozz37/glucose-measure-api/services/db"
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
	measure, err := models.GetLastMeasureFromCSV(fileName)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": "Can't get glucose level from CSV"})
	}

	db.Save(measure.GlucoseLevel, measure.Date)
	measure.Status = http.StatusOK
	c.JSON(measure.Status, measure)
}
