package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetStorages(c *gin.Context) {

	var storages []models.Storage
	result := database.DB.Find(&storages)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Storages"})
		return
	}

	c.JSON(200, storages)
}

func GetStorage(c *gin.Context) {

	storageID := c.Param("id")
	var storage models.Storage
	result := database.DB.First(&storage, storageID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Storage not found"})
		return
	}

	c.JSON(200, storage)
}

func GetStoragesByBrand(c *gin.Context) {

	storageBrand := c.Param("brand")
	var storages []models.Storage
	result := database.DB.Where("brand = ?", storageBrand).Find(&storages)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Storages by brand"})
		return
	}

	c.JSON(200, storages)
}

func CreateStorage(c *gin.Context) {

	var storage models.Storage

	if err := c.BindJSON(&storage); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&storage)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"storage": storage,
	})
}

func UpdateStorage(c *gin.Context) {

	storageID := c.Param("id")
	var storage models.Storage
	result := database.DB.First(&storage, storageID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Storage not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&storage).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Storage"})
		return
	}

	c.JSON(200, storage)
}

func DeleteStorage(c *gin.Context) {

	storageID := c.Param("id")
	var storage models.Storage
	result := database.DB.First(&storage, storageID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Storage not found"})
		return
	}

	result = database.DB.Delete(&storage)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Storage"})
		return
	}

	c.JSON(200, gin.H{"message": "Storage deleted successfully"})
}
