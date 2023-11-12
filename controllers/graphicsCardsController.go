package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetGraphicsCards(c *gin.Context) {

	var gpus []models.GraphicsCard
	result := database.DB.Find(&gpus)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Graphics Cards"})
		return
	}

	c.JSON(200, gpus)
}

func GetGraphicsCard(c *gin.Context) {

	gpuID := c.Param("id")
	var gpu models.GraphicsCard
	result := database.DB.First(&gpu, gpuID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Graphics Card not found"})
		return
	}

	c.JSON(200, gpu)
}

func GetGraphicsCardsByBrand(c *gin.Context) {

	gpuBrand := c.Param("brand")
	var gpu []models.GraphicsCard
	result := database.DB.Where("brand = ?", gpuBrand).Find(&gpu)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Graphics Cards by brand"})
		return
	}

	c.JSON(200, gpu)
}

func CreateGraphicsCard(c *gin.Context) {

	var gpu models.GraphicsCard

	if err := c.BindJSON(&gpu); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&gpu)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"graphicsCard": gpu,
	})
}

func UpdateGraphicsCard(c *gin.Context) {

	gpuID := c.Param("id")
	var gpu models.GraphicsCard
	result := database.DB.First(&gpu, gpuID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cooling System not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&gpu).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Graphics Card"})
		return
	}

	c.JSON(200, gpu)
}

func DeleteGraphicsCard(c *gin.Context) {

	gpuID := c.Param("id")
	var gpu models.GraphicsCard
	result := database.DB.First(&gpu, gpuID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Graphics Card not found"})
		return
	}

	result = database.DB.Delete(&gpu)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Graphics Card"})
		return
	}

	c.JSON(200, gin.H{"message": "Graphics Card deleted successfully"})
}
