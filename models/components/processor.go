package models

import "gorm.io/gorm"

type Processor struct {
	gorm.Model
	Brand              *string  `gorm:"not null;size:50"`
	ModelName          *string  `gorm:"not null;size:80"`
	SpeedGHz           *float64 `gorm:"not null"`
	Cores              *int     `gorm:"not null"`
	Threads            int
	SocketType         string
	IntegratedGraphics bool
	ImageURL           *string `gorm:"not null;size:200"`
}
