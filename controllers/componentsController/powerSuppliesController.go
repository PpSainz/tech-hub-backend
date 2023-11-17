package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetPowerSupplies(c *gin.Context) {

	var power []models.PowerSupply
	result := database.DB.Find(&power)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Power Supplies"})
		return
	}

	c.JSON(200, power)
}

func GetPowerSupply(c *gin.Context) {

	powerID := c.Param("id")
	var power models.PowerSupply
	result := database.DB.First(&power, powerID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Power Supply not found"})
		return
	}

	c.JSON(200, power)
}

func GetPowerSuppliesByBrand(c *gin.Context) {

	powerBrand := c.Param("brand")
	var power []models.PowerSupply
	result := database.DB.Where("brand = ?", powerBrand).Find(&power)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Power Supplies by brand"})
		return
	}

	c.JSON(200, power)
}

func CreatePowerSupply(c *gin.Context) {

	var power models.PowerSupply

	if err := c.BindJSON(&power); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&power)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"powerSupply": power,
	})
}

func UpdatePowerSupply(c *gin.Context) {

	powerID := c.Param("id")
	var power models.PowerSupply
	result := database.DB.First(&power, powerID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Power Supply not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&power).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Power Supply"})
		return
	}

	c.JSON(200, power)
}

func DeletePowerSupply(c *gin.Context) {

	powerID := c.Param("id")
	var power models.PowerSupply
	result := database.DB.First(&power, powerID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Power Supply not found"})
		return
	}

	result = database.DB.Delete(&power)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Power Supply"})
		return
	}

	c.JSON(200, gin.H{"message": "Power Supply deleted successfully"})
}
