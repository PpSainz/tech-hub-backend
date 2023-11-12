package models

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	Brand        *string `gorm:"not null;size:50"`
	ModelName    *string `gorm:"not null;size:80"`
	FormFactor   string  `gorm:"size:50"`
	NumDriveBays int
	ImageURL     *string `gorm:"not null;size:200"`
}
