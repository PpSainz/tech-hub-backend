package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func HelloTechHub(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ESTO ES TECH-HUB",
	})
}

func GetProcessors(c *gin.Context) {

	var processors []models.Processor
	result := database.DB.Find(&processors)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving processors"})
		return
	}

	c.JSON(200, processors)
}

func GetProcessor(c *gin.Context) {

	processorID := c.Param("id")
	var processor models.Processor
	result := database.DB.First(&processor, processorID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Processor not found"})
		return
	}

	c.JSON(200, processor)
}

func GetProcessorsByBrand(c *gin.Context) {

	processorBrand := c.Param("brand")
	var processors []models.Processor
	result := database.DB.Where("brand = ?", processorBrand).Find(&processors)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving processors by brand"})
		return
	}

	c.JSON(200, processors)
}

func CreateProcessor(c *gin.Context) {

	var processor models.Processor

	if err := c.BindJSON(&processor); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&processor)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"processor": processor,
	})
}

func UpdateProcessor(c *gin.Context) {

	processorID := c.Param("id")
	var processor models.Processor
	result := database.DB.First(&processor, processorID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Processor not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&processor).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating processor"})
		return
	}

	c.JSON(200, processor)
}

func DeleteProcessor(c *gin.Context) {

	processorID := c.Param("id")
	var processor models.Processor
	result := database.DB.First(&processor, processorID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Processor not found"})
		return
	}

	result = database.DB.Delete(&processor)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting processor"})
		return
	}

	c.JSON(200, gin.H{"message": "Processor deleted successfully"})
}
