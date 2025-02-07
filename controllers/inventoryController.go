package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek-srivastav-omniful/wms-service/models"
)

func CreateInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	inventory.CreatedAt = time.Now()
	inventory.UpdatedAt = time.Now()

	if err := DB.Create(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inventory created successfully", "inventory": inventory})
}

func GetInventoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var inventory models.Inventory
	if err := DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory record not found"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func GetAllInventory(c *gin.Context) {
	var inventories []models.Inventory
	if err := DB.Find(&inventories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory records"})
		return
	}
	c.JSON(http.StatusOK, inventories)
}

func UpdateInventory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	inventory.UpdatedAt = time.Now()

	if err := DB.Model(&models.Inventory{}).Where("id = ?", id).Updates(inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully", "inventory": inventory})
}

func DeleteInventory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := DB.Delete(&models.Inventory{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully"})
}
