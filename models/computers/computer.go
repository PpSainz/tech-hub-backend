package models

import (
	models "github.com/Yiro13/tech-hub-backend/models/components"
	"gorm.io/gorm"
)

type Computer struct {
	gorm.Model
	Brand           string             `gorm:"size:50"`
	ModelName       string             `gorm:"size:80"`
	ProcessorID     *uint              `gorm:"not null"`
	Processor       models.Processor   `gorm:"foreignKey:ProcessorID"`
	MotherBoardID   *uint              `gorm:"not null"`
	MotherBoard     models.MotherBoard `gorm:"foreignKey:MotherBoardID"`
	GraphicsCardID  uint
	GraphicsCard    models.GraphicsCard `gorm:"foreignKey:GraphicsCardID"`
	RamMemorieID    *uint               `gorm:"not null"`
	RamMemorie      models.RAM          `gorm:"foreignKey:RamMemorieID"`
	StorageID       *uint               `gorm:"not null"`
	Storage         models.Storage      `gorm:"foreignKey:StorageID"`
	PowerSupplyID   *uint               `gorm:"not null"`
	PowerSupply     models.PowerSupply  `gorm:"foreignKey:PowerSupplyID"`
	CaseID          *uint               `gorm:"not null"`
	Case            models.Case         `gorm:"foreignKey:CaseID"`
	CoolingSystemID uint
	CoolingSystem   models.CoolingSystem `gorm:"foreignKey:CoolingSystemID"`
	OSID            *uint                `gorm:"not null"`
	OS              models.OS            `gorm:"foreignKey:OSID"`
}
