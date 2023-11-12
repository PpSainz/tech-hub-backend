package main

import (
	"github.com/Yiro13/tech-hub-backend/controllers"
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"github.com/gin-gonic/gin"
)

func main() {

	database.DBConnection()
	database.DB.AutoMigrate(&models.Processor{})
	database.DB.AutoMigrate(&models.MotherBoard{})

	r := gin.Default()
	r.GET("/", controllers.HelloTechHub)
	r.GET("/processors", controllers.GetProcessors)
	r.GET("/processors/:id", controllers.GetProcessor)
	r.GET("/processors/brands/:brand", controllers.GetProcessorsByBrand)
	r.POST("/processors", controllers.CreateProcessor)
	r.PUT("/processors/:id", controllers.UpdateProcessor)
	r.DELETE("/processors/:id", controllers.DeleteProcessor)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
