package main

import (
	"github.com/Yiro13/tech-hub-backend/database"
	"github.com/Yiro13/tech-hub-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {

	database.DBConnection()
	database.DB.AutoMigrate(&models.Software{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hola mundooo")
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}