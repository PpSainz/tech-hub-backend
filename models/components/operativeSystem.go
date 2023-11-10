package models

import "gorm.io/gorm"

type OS struct {
	gorm.Model
	Brand       string `gorm:"not null;size:50"`
	Name        string `gorm:"not null;size:50"`
	Version     string `gorm:"not null;size:10"`
	Description string `gorm:"not null;size:200"`
}
