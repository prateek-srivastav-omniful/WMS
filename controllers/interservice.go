package controllers

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prateek-srivastav-omniful/wms-service/models"
	"gorm.io/gorm"
)

func ValidateSKU(c *gin.Context) {
	skuID := c.Param("id")
	fmt.Println("validate sku called on id -> ", skuID)

	skuIDInt, err := strconv.Atoi(skuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sku_id"})
		return
	}

	// Check if the SKU exists in the database
	var sku models.SKU
	err = DB.Where("id = ?", skuIDInt).First(&sku).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check SKU existence"})
		}
		return
	}

	// If SKU exists, return success message
	c.JSON(http.StatusOK, gin.H{"message": "SKU exists"})
}

func ValidateHub(c *gin.Context) {
	hubID := c.Param("id")

	// Convert to proper data type
	hubIDInt, err := strconv.Atoi(hubID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub_id"})
		return
	}

	// Check if the Hub exists in the database
	var hub models.Hub
	err = DB.Where("id = ?", hubIDInt).First(&hub).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check hub existence"})
		}
		return
	}

	// If hub exists, return success message
	c.JSON(http.StatusOK, gin.H{"message": "Hub exists"})
}
