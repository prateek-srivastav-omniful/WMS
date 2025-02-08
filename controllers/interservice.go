package controllers

import (
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prateek-srivastav-omniful/wms-service/models"
	"gorm.io/gorm"
)

func ValidateSKU(c *gin.Context) {
	skuID := c.Param("skuid")

	log.Println("validate sku called on id -> ", skuID)
	c.JSON(http.StatusOK, gin.H{"message": "SKU exists"})
	return
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

func ValidateInventorySKU(c *gin.Context) {
	skuID := c.Param("skuid")
	log.Println("validate sku called on id -> ", skuID)
	c.JSON(http.StatusOK, gin.H{"message": "SKU Quantity>0"})
	return
	skuIDInt, err := strconv.Atoi(skuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
		return
	}

	var inventory models.Inventory
	err = DB.Where("sku_id = ?", skuIDInt).First(&inventory).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found in inventory"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check inventory"})
		}
		return
	}

	if inventory.Quantity <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU out of stock"})
		return
	}

	// Decrement inventory quantity
	err = DB.Model(&inventory).Where("sku_id = ?", skuIDInt).UpdateColumn("quantity", gorm.Expr("quantity - 1")).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SKU exists, in stock, and quantity decremented"})
}
func ValidateHub(c *gin.Context) {

}
