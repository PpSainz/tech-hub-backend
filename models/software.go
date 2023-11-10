package models

import "gorm.io/gorm"

type Software struct {
	gorm.Model
	Name        string `gorm:"not null;size:50"`
	Version     string `gorm:"not null;size:10"`
	Description string `gorm:"not null;size:200"`
}
