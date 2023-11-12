package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetCoolingSystems(c *gin.Context) {

	var coolings []models.CoolingSystem
	result := database.DB.Find(&coolings)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Cooling Systems"})
		return
	}

	c.JSON(200, coolings)
}

func GetCollingSystem(c *gin.Context) {

	coolingID := c.Param("id")
	var cooling models.CoolingSystem
	result := database.DB.First(&cooling, coolingID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cooling System not found"})
		return
	}

	c.JSON(200, cooling)
}

func GetCollingSystemByBrand(c *gin.Context) {

	coolingBrand := c.Param("brand")
	var cooling []models.CoolingSystem
	result := database.DB.Where("brand = ?", coolingBrand).Find(&cooling)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Cooling Systems by brand"})
		return
	}

	c.JSON(200, cooling)
}

func CreateCollingSystem(c *gin.Context) {

	var cooling models.CoolingSystem

	if err := c.BindJSON(&cooling); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&cooling)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"coolingSystem": cooling,
	})
}

func UpdateCoolingSystem(c *gin.Context) {

	coolingID := c.Param("id")
	var cooling models.CoolingSystem
	result := database.DB.First(&cooling, coolingID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cooling System not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&cooling).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Cooling System"})
		return
	}

	c.JSON(200, cooling)
}

func DeleteCoolingSystem(c *gin.Context) {

	coolingID := c.Param("id")
	var cooling models.CoolingSystem
	result := database.DB.First(&cooling, coolingID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cooling System not found"})
		return
	}

	result = database.DB.Delete(&cooling)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Cooling System"})
		return
	}

	c.JSON(200, gin.H{"message": "Cooling System deleted successfully"})
}
