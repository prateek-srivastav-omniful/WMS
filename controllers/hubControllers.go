package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prateek-srivastav-omniful/wms-service/models"
	"gorm.io/gorm"
)

// Assume we have a global DB instance (replace with actual DB instance)
var DB *gorm.DB

func SetDbInstance(db *gorm.DB) {
	DB = db
}

// CreateHub handles the creation of a new Hub
func CreateHub(c *gin.Context) {
	var hub models.Hub

	// Bind JSON body to the Hub struct
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Save to DB
	if err := DB.Create(&hub).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hub"})
		return
	}

	c.JSON(http.StatusCreated, hub)
}

// GetHubs retrieves all hubs from the database
func GetHubs(c *gin.Context) {
	var hubs []models.Hub

	// Fetch all hubs
	if err := DB.Find(&hubs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hubs"})
		return
	}

	c.JSON(http.StatusOK, hubs)
}
