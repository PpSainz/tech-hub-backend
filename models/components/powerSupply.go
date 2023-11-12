package models

import "gorm.io/gorm"

type PowerSupply struct {
	gorm.Model
	Brand         *string `gorm:"not null;size:50"`
	ModelName     *string `gorm:"not null;size:80"`
	Wattage       *int    `gorm:"not null"`
	Certification *string `gorm:"not null;size:80"`
	ImageURL      *string `gorm:"not null;size:200"`
}
