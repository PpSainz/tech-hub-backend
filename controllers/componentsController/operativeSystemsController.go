package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetOSs(c *gin.Context) {

	var oss []models.OS
	result := database.DB.Find(&oss)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Operative Systems"})
		return
	}

	c.JSON(200, oss)
}

func GetOS(c *gin.Context) {

	osID := c.Param("id")
	var os models.OS
	result := database.DB.First(&os, osID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Operative System not found"})
		return
	}

	c.JSON(200, os)
}

func GetOSByBrand(c *gin.Context) {

	osBrand := c.Param("brand")
	var os []models.OS
	result := database.DB.Where("brand = ?", osBrand).Find(&os)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Operative Systems by brand"})
		return
	}

	c.JSON(200, os)
}

func CreateOS(c *gin.Context) {

	var os models.OS

	if err := c.BindJSON(&os); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&os)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"operativeSystem": os,
	})
}

func UpdateOS(c *gin.Context) {

	osID := c.Param("id")
	var os models.OS
	result := database.DB.First(&os, osID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Operative System not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&os); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&os).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Graphics Card"})
		return
	}

	c.JSON(200, os)
}

func DeleteOS(c *gin.Context) {

	osID := c.Param("id")
	var os models.OS
	result := database.DB.First(&os, osID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Operative System not found"})
		return
	}

	result = database.DB.Delete(&os)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Operative System"})
		return
	}

	c.JSON(200, gin.H{"message": "Operative System deleted successfully"})
}
