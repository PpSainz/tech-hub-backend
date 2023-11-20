package routes

import (
	controllers "github.com/Yiro13/tech-hub-backend/controllers/componentsController"
	pcscontroller "github.com/Yiro13/tech-hub-backend/controllers/pcsController"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/", controllers.HelloTechHub)

	r.GET("/processors", controllers.GetProcessors)
	r.GET("/mobos", controllers.GetMotherBoards)
	r.GET("/ramMemories", controllers.GetRams)
	r.GET("/cases", controllers.GetCases)
	r.GET("/coolings", controllers.GetCoolingSystems)
	r.GET("/gpus", controllers.GetGraphicsCards)
	r.GET("/os", controllers.GetOSs)
	r.GET("/power", controllers.GetPowerSupplies)
	r.GET("/storages", controllers.GetStorages)
	r.GET("/computers", pcscontroller.GetComputers)
	r.GET("/recommendations", pcscontroller.GetRecommendations)

	r.GET("/processors/:id", controllers.GetProcessor)
	r.GET("/mobos/:id", controllers.GetMotherBoard)
	r.GET("/ramMemories/:id", controllers.GetRam)
	r.GET("/cases/:id", controllers.GetCase)
	r.GET("/coolings/:id", controllers.GetCollingSystem)
	r.GET("/gpus/:id", controllers.GetGraphicsCard)
	r.GET("/os/:id", controllers.GetOS)
	r.GET("/power/:id", controllers.GetPowerSupply)
	r.GET("/storages/:id", controllers.GetStorage)
	r.GET("/computers/:id", pcscontroller.GetComputer)
	r.GET("/recommendations/:id", pcscontroller.GetRecommendation)

	r.GET("/processors/brands/:brand", controllers.GetProcessorsByBrand)
	r.GET("/mobos/brands/:brand", controllers.GetMoboByBrand)
	r.GET("/ramMemories/brands/:brand", controllers.GetRamByBrand)
	r.GET("/cases/brands/:brand", controllers.GetCaseByBrand)
	r.GET("/coolings/brands/:brand", controllers.GetCollingSystemByBrand)
	r.GET("/gpus/brands/:brand", controllers.GetGraphicsCardsByBrand)
	r.GET("/os/brands/:brand", controllers.GetOSByBrand)
	r.GET("/power/brands/:brand", controllers.GetPowerSuppliesByBrand)
	r.GET("/storages/brands/:brand", controllers.GetStoragesByBrand)
	r.GET("/computers/brands/:brand", pcscontroller.GetComputerByBrand)

	r.POST("/processors", controllers.CreateProcessor)
	r.POST("/mobos", controllers.CreateMotherBoard)
	r.POST("/ramMemories", controllers.CreateRam)
	r.POST("/cases", controllers.CreateCase)
	r.POST("/coolings", controllers.CreateCollingSystem)
	r.POST("/gpus", controllers.CreateGraphicsCard)
	r.POST("/os", controllers.CreateOS)
	r.POST("/power", controllers.CreatePowerSupply)
	r.POST("/storages", controllers.CreateStorage)
	r.POST("/computers", pcscontroller.CreateComputer)
	r.POST("/recommendations", pcscontroller.CreateRecommendation)

	r.PUT("/processors/:id", controllers.UpdateProcessor)
	r.PUT("/mobos/:id", controllers.UpdateMotherBoard)
	r.PUT("/ramMemories/:id", controllers.UpdateRam)
	r.PUT("/cases/:id", controllers.UpdateCase)
	r.PUT("/coolings/:id", controllers.UpdateCoolingSystem)
	r.PUT("/gpus/:id", controllers.UpdateGraphicsCard)
	r.PUT("/os/:id", controllers.UpdateOS)
	r.PUT("/power/:id", controllers.UpdatePowerSupply)
	r.PUT("/storages/:id", controllers.UpdateStorage)
	r.PUT("/computers/:id", pcscontroller.UpdateComputer)
	r.PUT("/recommendations/:id", pcscontroller.UpdateRecommendation)

	r.DELETE("/processors/:id", controllers.DeleteProcessor)
	r.DELETE("/mobos/:id", controllers.DeleteMotherBoard)
	r.DELETE("/ramMemories/:id", controllers.DeleteRam)
	r.DELETE("/cases/:id", controllers.DeleteCase)
	r.DELETE("/coolings/:id", controllers.DeleteCoolingSystem)
	r.DELETE("/gpus/:id", controllers.DeleteGraphicsCard)
	r.DELETE("/os/:id", controllers.DeleteOS)
	r.DELETE("/power/:id", controllers.DeletePowerSupply)
	r.DELETE("/storages/:id", controllers.DeleteStorage)
	r.DELETE("/computers/:id", pcscontroller.DeleteComputer)
	r.DELETE("/recommendations/:id", pcscontroller.DeleteRecommendation)

	r.GET("/mobos/socket/:sockettype", controllers.GetMoBoBySocket)
	r.GET("/computers/models/:modelName", pcscontroller.GetComputerByModelName)
	r.GET("/recommendations/tags/:tag", pcscontroller.GetRecommendationByTag)

	return r
}
