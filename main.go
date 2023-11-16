package main

import (
	"github.com/Yiro13/tech-hub-backend/controllers"
	"github.com/Yiro13/tech-hub-backend/database"
	models "github.com/Yiro13/tech-hub-backend/models/components"
	computers "github.com/Yiro13/tech-hub-backend/models/computers"
	"github.com/gin-contrib/cors"
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
	database.DB.AutoMigrate(&models.PowerSupply{})
	database.DB.AutoMigrate(&models.Storage{})
	database.DB.AutoMigrate(&computers.Computer{})

	r := gin.Default()
	r.GET("/", controllers.HelloTechHub)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/processors", controllers.GetProcessors)
	r.GET("/mobos", controllers.GetMotherBoards)
	r.GET("/ramMemories", controllers.GetRams)
	r.GET("/cases", controllers.GetCases)
	r.GET("/coolings", controllers.GetCoolingSystems)
	r.GET("/gpus", controllers.GetGraphicsCards)
	r.GET("/os", controllers.GetOSs)
	r.GET("/power", controllers.GetPowerSupplies)
	r.GET("/storages", controllers.GetStorages)

	r.GET("/processors/:id", controllers.GetProcessor)
	r.GET("/mobos/:id", controllers.GetMotherBoard)
	r.GET("/ramMemories/:id", controllers.GetRam)
	r.GET("/cases/:id", controllers.GetCase)
	r.GET("/coolings/:id", controllers.GetCollingSystem)
	r.GET("/gpus/:id", controllers.GetGraphicsCard)
	r.GET("/os/:id", controllers.GetOS)
	r.GET("/power/:id", controllers.GetPowerSupply)
	r.GET("/storages/:id", controllers.GetStorage)

	r.GET("/processors/brands/:brand", controllers.GetProcessorsByBrand)
	r.GET("/mobos/brands/:brand", controllers.GetMoboByBrand)
	r.GET("/ramMemories/brands/:brand", controllers.GetRamByBrand)
	r.GET("/cases/brands/:brand", controllers.GetCaseByBrand)
	r.GET("/coolings/brands/:brand", controllers.GetCollingSystemByBrand)
	r.GET("/gpus/brands/:brand", controllers.GetGraphicsCardsByBrand)
	r.GET("/os/brands/:brand", controllers.GetOSByBrand)
	r.GET("/power/brands/:brand", controllers.GetPowerSuppliesByBrand)
	r.GET("/storages/brands/:brand", controllers.GetStoragesByBrand)

	r.POST("/processors", controllers.CreateProcessor)
	r.POST("/mobos", controllers.CreateMotherBoard)
	r.POST("/ramMemories", controllers.CreateRam)
	r.POST("/cases", controllers.CreateCase)
	r.POST("/coolings", controllers.CreateCollingSystem)
	r.POST("/gpus", controllers.CreateGraphicsCard)
	r.POST("/os", controllers.CreateOS)
	r.POST("/power", controllers.CreatePowerSupply)
	r.POST("/storages", controllers.CreateStorage)

	r.PUT("/processors/:id", controllers.UpdateProcessor)
	r.PUT("/mobos/:id", controllers.UpdateMotherBoard)
	r.PUT("/ramMemories/:id", controllers.UpdateRam)
	r.PUT("/cases/:id", controllers.UpdateCase)
	r.PUT("/coolings/:id", controllers.UpdateCoolingSystem)
	r.PUT("/gpus/:id", controllers.UpdateGraphicsCard)
	r.PUT("/os/:id", controllers.UpdateOS)
	r.PUT("/power", controllers.UpdatePowerSupply)
	r.PUT("/storages", controllers.UpdateStorage)

	r.DELETE("/processors/:id", controllers.DeleteProcessor)
	r.DELETE("/mobos/:id", controllers.DeleteMotherBoard)
	r.DELETE("/ramMemories/:id", controllers.DeleteRam)
	r.DELETE("/cases/:id", controllers.DeleteCase)
	r.DELETE("/coolings/:id", controllers.DeleteCoolingSystem)
	r.DELETE("/gpus/:id", controllers.DeleteGraphicsCard)
	r.DELETE("/os/:id", controllers.DeleteOS)
	r.DELETE("/power/:id", controllers.DeletePowerSupply)
	r.DELETE("/storages/:id", controllers.DeleteStorage)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
