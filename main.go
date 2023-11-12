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
	database.DB.AutoMigrate(&models.RAM{})
	database.DB.AutoMigrate(&models.Case{})

	r := gin.Default()
	r.GET("/", controllers.HelloTechHub)

	r.GET("/processors", controllers.GetProcessors)
	r.GET("/mobos", controllers.GetMotherBoards)
	r.GET("/ramMemories", controllers.GetRams)
	r.GET("/cases", controllers.GetCases)

	r.GET("/processors/:id", controllers.GetProcessor)
	r.GET("/mobos/:id", controllers.GetMotherBoard)
	r.GET("/ramMemories/:id", controllers.GetRam)
	r.GET("/cases/:id", controllers.GetCase)

	r.GET("/processors/brands/:brand", controllers.GetProcessorsByBrand)
	r.GET("/mobos/brands/:brand", controllers.GetMoboByBrand)
	r.GET("/ramMemories/brands/:brand", controllers.GetRamByBrand)
	r.GET("/cases/brands/:brand", controllers.GetCaseByBrand)

	r.POST("/processors", controllers.CreateProcessor)
	r.POST("/mobos", controllers.CreateMotherBoard)
	r.POST("/ramMemories", controllers.CreateRam)
	r.POST("/cases", controllers.CreateCase)

	r.PUT("/processors/:id", controllers.UpdateProcessor)
	r.PUT("/mobos/:id", controllers.UpdateMotherBoard)
	r.PUT("/ramMemories/:id", controllers.UpdateRam)
	r.PUT("/cases/:id", controllers.UpdateCase)

	r.DELETE("/processors/:id", controllers.DeleteProcessor)
	r.DELETE("/mobos/:id", controllers.DeleteMotherBoard)
	r.DELETE("/ramMemories/:id", controllers.DeleteRam)
	r.DELETE("/cases/:id", controllers.DeleteCase)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
