package pcscontroller

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/computers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetComputers(c *gin.Context) {

	var computers []models.Computer
	result := database.DB.Find(&computers)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Computers"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&computers)

	c.JSON(200, computers)
}

func GetComputer(c *gin.Context) {

	computerID := c.Param("id")
	var computer models.Computer
	result := database.DB.First(&computer, computerID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Computer not found"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&computer)

	c.JSON(200, computer)

}

func GetComputerByBrand(c *gin.Context) {

	computerBrand := c.Param("brand")
	var computers []models.Computer
	result := database.DB.Where("brand = ?", computerBrand).Find(&computers)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Computers by brand"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&computers)

	c.JSON(200, computers)
}

func GetComputerByModelName(c *gin.Context) {

	computerModel := c.Param("modelName")
	var computers []models.Computer
	result := database.DB.Where("model_name = ?", computerModel).Find(&computers)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Computers by name model"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&computers)

	c.JSON(200, computers)
}

func CreateComputer(c *gin.Context) {

	var computer models.Computer

	if err := c.BindJSON(&computer); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if computer.Brand == "" {
		computer.Brand = "Custom"
	}

	if computer.ModelName == "" {
		computer.ModelName = "Custom Model"
	}

	result := database.DB.Create(&computer)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	database.DB.Preload(clause.Associations).Find(&computer)

	c.JSON(201, gin.H{
		"computer": computer,
	})
}

func UpdateComputer(c *gin.Context) {

	computerID := c.Param("id")
	var computer models.Computer
	result := database.DB.First(&computer, computerID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Computer not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&computer).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Computer"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&computer)

	c.JSON(200, computer)
}

func DeleteComputer(c *gin.Context) {

	computerID := c.Param("id")
	var computer models.Computer
	result := database.DB.First(&computer, computerID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Computer not found"})
		return
	}

	result = database.DB.Delete(&computer)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Computer"})
		return
	}

	c.JSON(200, gin.H{"message": "Computer deleted successfully"})
}
