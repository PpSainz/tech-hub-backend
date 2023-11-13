package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	Brand         *string `gorm:"not null;size:50"`
	ModelName     *string `gorm:"not null;size:80"`
	CapacityGB    *int    `gorm:"not null"`
	Type          *string `gorm:"not null"`
	Interface     *string `gorm:"not null;size:80"`
	RotationSpeed int
	ImageURL      *string `gorm:"not null;size:200"`
}
