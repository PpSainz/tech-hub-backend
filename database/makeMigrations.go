package database

import (
	models "github.com/Yiro13/tech-hub-backend/models/components"
	computers "github.com/Yiro13/tech-hub-backend/models/computers"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Processor{}, &models.MotherBoard{}, &models.RAM{}, &models.Case{}, &models.CoolingSystem{}, &models.GraphicsCard{}, &models.OS{},
		&models.PowerSupply{}, &models.Storage{}, &computers.Computer{}, &computers.Recommendation{})
}
