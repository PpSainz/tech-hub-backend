package controllers

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func GetCases(c *gin.Context) {

	var cases []models.Case
	result := database.DB.Find(&cases)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Cases"})
		return
	}

	c.JSON(200, cases)
}

func GetCase(c *gin.Context) {

	caseID := c.Param("id")
	var cased models.Case
	result := database.DB.First(&cased, caseID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Case not found"})
		return
	}

	c.JSON(200, cased)
}

func GetCaseByBrand(c *gin.Context) {

	caseBrand := c.Param("brand")
	var cases []models.Case
	result := database.DB.Where("brand = ?", caseBrand).Find(&cases)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Cases by brand"})
		return
	}

	c.JSON(200, cases)
}

func CreateCase(c *gin.Context) {

	var cased models.Case

	if err := c.BindJSON(&cased); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&cased)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"case": cased,
	})
}

func UpdateCase(c *gin.Context) {

	caseID := c.Param("id")
	var cased models.Case
	result := database.DB.First(&cased, caseID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Case not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&cased).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating MotherBoard"})
		return
	}

	c.JSON(200, cased)
}

func DeleteCase(c *gin.Context) {

	caseID := c.Param("id")
	var cased models.Case
	result := database.DB.First(&cased, caseID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Case not found"})
		return
	}

	result = database.DB.Delete(&cased)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Case"})
		return
	}

	c.JSON(200, gin.H{"message": "Case deleted successfully"})
}
