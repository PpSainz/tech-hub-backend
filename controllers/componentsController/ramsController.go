package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetRams(c *gin.Context) {

	var rams []models.RAM
	result := database.DB.Find(&rams)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving RAM Memories"})
		return
	}

	c.JSON(200, rams)
}

func GetRam(c *gin.Context) {

	ramID := c.Param("id")
	var ram models.RAM
	result := database.DB.First(&ram, ramID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "RAM Memory not found"})
		return
	}

	c.JSON(200, ram)
}

func GetRamByBrand(c *gin.Context) {

	ramBrand := c.Param("brand")
	var rams []models.RAM
	result := database.DB.Where("brand = ?", ramBrand).Find(&rams)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving RAM Memories by brand"})
		return
	}

	c.JSON(200, rams)
}

func CreateRam(c *gin.Context) {

	var ram models.RAM

	if err := c.BindJSON(&ram); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&ram)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"ram": ram,
	})
}

func UpdateRam(c *gin.Context) {

	ramID := c.Param("id")
	var ram models.RAM
	result := database.DB.First(&ram, ramID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "RAM Memory not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&ram).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating RAM Memory"})
		return
	}

	c.JSON(200, ram)
}

func DeleteRam(c *gin.Context) {

	ramID := c.Param("id")
	var ram models.RAM
	result := database.DB.First(&ram, ramID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "RAM Memory not found"})
		return
	}

	result = database.DB.Delete(&ram)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting RAM Memory"})
		return
	}

	c.JSON(200, gin.H{"message": "RAM Memory deleted successfully"})
}
