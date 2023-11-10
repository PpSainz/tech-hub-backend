package main

import (
	"github.com/Yiro13/tech-hub-backend/database"
	//models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func main() {

	database.DBConnection()
	//database.DB.AutoMigrate(&models.CoolingSystem{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hola munde")
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
