package pcscontroller

import (
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/computers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetRecommendations(c *gin.Context) {

	var recommendations []models.Recommendation
	result := database.DB.Find(&recommendations)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error retrieving Recommendations"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&recommendations)
	for i := 0; i < len(recommendations); i++ {
		database.DB.Preload(clause.Associations).Find(&recommendations[i].Computer)
	}

	c.JSON(200, recommendations)
}

func GetRecommendation(c *gin.Context) {

	recommendationID := c.Param("id")
	var recommendation models.Recommendation
	result := database.DB.First(&recommendation, recommendationID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Recommendation not found"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&recommendation.Computer)

	c.JSON(200, recommendation)
}

func GetRecommendationByTag(c *gin.Context) {

	recommendationTag := c.Param("tag")
	var recommendations []models.Recommendation
	result := database.DB.Where("tag = ?", recommendationTag).Find(&recommendations)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Error retrieving Recomendations by tag"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&recommendations)
	for i := 0; i < len(recommendations); i++ {
		database.DB.Preload(clause.Associations).Find(&recommendations[i].Computer)
	}

	c.JSON(200, recommendations)
}

func CreateRecommendation(c *gin.Context) {

	var recommendation models.Recommendation

	if err := c.BindJSON(&recommendation); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Create(&recommendation)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	database.DB.Preload(clause.Associations).Find(&recommendation.Computer)

	c.JSON(201, gin.H{
		"recommendation": recommendation,
	})
}

func UpdateRecommendation(c *gin.Context) {

	recommendationID := c.Param("id")
	var recommendation models.Recommendation
	result := database.DB.First(&recommendation, recommendationID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Recommendation not found"})
		return
	}

	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	result = database.DB.Model(&recommendation).Updates(updateData)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error updating Recommendation"})
		return
	}

	database.DB.Preload(clause.Associations).Find(&recommendation.Computer)

	c.JSON(200, recommendation)
}

func DeleteRecommendation(c *gin.Context) {

	recommendationID := c.Param("id")
	var recommendation models.Recommendation
	result := database.DB.First(&recommendation, recommendationID)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Recommendation not found"})
		return
	}

	result = database.DB.Delete(&recommendation)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting Recommendation"})
		return
	}

	c.JSON(200, gin.H{"message": "Recommendation deleted successfully"})
}
