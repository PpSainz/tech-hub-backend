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
	database.DB.AutoMigrate(&models.CoolingSystem{})
	database.DB.AutoMigrate(&models.GraphicsCard{})
	database.DB.AutoMigrate(&models.OS{})

	r := gin.Default()
	r.GET("/", controllers.HelloTechHub)

	r.GET("/processors", controllers.GetProcessors)
	r.GET("/mobos", controllers.GetMotherBoards)
	r.GET("/ramMemories", controllers.GetRams)
	r.GET("/cases", controllers.GetCases)
	r.GET("/coolings", controllers.GetCoolingSystems)
	r.GET("/gpus", controllers.GetGraphicsCards)
	r.GET("/os", controllers.GetOSs)

	r.GET("/processors/:id", controllers.GetProcessor)
	r.GET("/mobos/:id", controllers.GetMotherBoard)
	r.GET("/ramMemories/:id", controllers.GetRam)
	r.GET("/cases/:id", controllers.GetCase)
	r.GET("/coolings/:id", controllers.GetCollingSystem)
	r.GET("/gpus/:id", controllers.GetGraphicsCard)
	r.GET("/os/:id", controllers.GetOS)

	r.GET("/processors/brands/:brand", controllers.GetProcessorsByBrand)
	r.GET("/mobos/brands/:brand", controllers.GetMoboByBrand)
	r.GET("/ramMemories/brands/:brand", controllers.GetRamByBrand)
	r.GET("/cases/brands/:brand", controllers.GetCaseByBrand)
	r.GET("/coolings/brands/:brand", controllers.GetCollingSystemByBrand)
	r.GET("/gpus/brands/:brand", controllers.GetGraphicsCardsByBrand)
	r.GET("/os/brands/:brand", controllers.GetOSByBrand)

	r.POST("/processors", controllers.CreateProcessor)
	r.POST("/mobos", controllers.CreateMotherBoard)
	r.POST("/ramMemories", controllers.CreateRam)
	r.POST("/cases", controllers.CreateCase)
	r.POST("/coolings", controllers.CreateCollingSystem)
	r.POST("/gpus", controllers.CreateGraphicsCard)
	r.POST("/os", controllers.CreateOS)

	r.PUT("/processors/:id", controllers.UpdateProcessor)
	r.PUT("/mobos/:id", controllers.UpdateMotherBoard)
	r.PUT("/ramMemories/:id", controllers.UpdateRam)
	r.PUT("/cases/:id", controllers.UpdateCase)
	r.PUT("/coolings/:id", controllers.UpdateCoolingSystem)
	r.PUT("/gpus/:id", controllers.UpdateGraphicsCard)
	r.PUT("/os/:id", controllers.UpdateOS)

	r.DELETE("/processors/:id", controllers.DeleteProcessor)
	r.DELETE("/mobos/:id", controllers.DeleteMotherBoard)
	r.DELETE("/ramMemories/:id", controllers.DeleteRam)
	r.DELETE("/cases/:id", controllers.DeleteCase)
	r.DELETE("/coolings/:id", controllers.DeleteCoolingSystem)
	r.DELETE("/gpus/:id", controllers.DeleteGraphicsCard)
	r.DELETE("/os/:id", controllers.UpdateOS)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
