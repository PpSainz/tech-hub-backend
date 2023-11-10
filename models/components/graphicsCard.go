package models

import "gorm.io/gorm"

type GraphicsCard struct {
	gorm.Model
	Brand     string `gorm:"not null;size:50"`
	ModelName string `gorm:"not null;size:80"`
	MemoryGB  int    `gorm:"not null"`
	Ports     int
	Interface string  `gorm:"not null;size:80"`
	SpeedGHz  float64 `gorm:"not null"`
	BoostGHz  float64
	ImageURL  string `gorm:"not null;size:200"`
}
