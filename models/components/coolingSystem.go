package models

import "gorm.io/gorm"

type CoolingSystem struct {
	gorm.Model
	Brand        string `gorm:"not null;size:50"`
	ModelName    string `gorm:"not null;size:80"`
	Type         string
	RadiatorSize float32
	FanSize      float32
	PumpIncluded bool   `gorm:"default:false"`
	ImageURL     string `gorm:"not null;size:200"`
}
