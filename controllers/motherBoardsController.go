package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetMotherBoards(c *gin.Context) {

	var motherBoards []models.MotherBoard
	result := database.DB.Find(&motherBoards)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Motherboards"})
		return
	}

	c.JSON(200, motherBoards)
}

func GetMotherBoard(c *gin.Context) {

	moboID := c.Param("id")
	var motherBoard models.MotherBoard
	result := database.DB.First(&motherBoard, moboID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "MotherBoard not found"})
		return
	}

	c.JSON(200, motherBoard)
}

func GetMoboByBrand(c *gin.Context) {

	moboBrand := c.Param("brand")
	var motherBoards []models.MotherBoard
	result := database.DB.Where("brand = ?", moboBrand).Find(&motherBoards)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving MotherBoards by brand"})
		return
	}

	c.JSON(200, motherBoards)
}

func CreateMotherBoard(c *gin.Context) {

	var motherBoard models.MotherBoard

	if err := c.BindJSON(&motherBoard); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&motherBoard)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"processor": motherBoard,
	})
}

func UpdateMotherBoard(c *gin.Context) {

	moboID := c.Param("id")
	var motherBoard models.MotherBoard
	result := database.DB.First(&motherBoard, moboID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "MotherBoard not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&motherBoard).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating MotherBoard"})
		return
	}

	c.JSON(200, motherBoard)
}

func DeleteMotherBoard(c *gin.Context) {

	moboID := c.Param("id")
	var motherBoard models.MotherBoard
	result := database.DB.First(&motherBoard, moboID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "MotherBoard not found"})
		return
	}

	result = database.DB.Delete(&motherBoard)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting MotherBoard"})
		return
	}

	c.JSON(200, gin.H{"message": "MotherBoard deleted successfully"})
}
