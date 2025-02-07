package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prateek-srivastav-omniful/wms-service/models"
)

// CreateSKU handles the creation of an SKU
func CreateSKU(c *gin.Context) {
	var sku models.SKU
	if err := c.ShouldBindJSON(&sku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := DB.Create(&sku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SKU"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "SKU created successfully", "sku": sku})
}

// GetSKUByID retrieves an SKU by ID
func GetSKUByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sku models.SKU
	if err := DB.First(&sku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}

	c.JSON(http.StatusOK, sku)
}

// GetAllSKUs retrieves all SKUs
func GetAllSKUs(c *gin.Context) {
	var skus []models.SKU
	if err := DB.Find(&skus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve SKUs"})
		return
	}
	c.JSON(http.StatusOK, skus)
}

// UpdateSKU updates an SKU
func UpdateSKU(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sku models.SKU
	if err := c.ShouldBindJSON(&sku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := DB.Model(&models.SKU{}).Where("id = ?", id).Updates(sku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update SKU"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SKU updated successfully", "sku": sku})
}

// DeleteSKU deletes an SKU by ID
func DeleteSKU(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := DB.Delete(&models.SKU{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete SKU"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SKU deleted successfully"})
}
