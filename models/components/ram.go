package models

import "gorm.io/gorm"

type RAM struct {
	gorm.Model
	Brand      string  `gorm:"not null;size:50"`
	ModelName  string  `gorm:"not null;size:80"`
	CapacityGB int     `gorm:"not null"`
	Type       string  `gorm:"not null"`
	SpeedGHz   float64 `gorm:"not null"`
	Modules    int
	ImageURL   string `gorm:"not null;size:200"`
}
