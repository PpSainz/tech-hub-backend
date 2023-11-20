package models

import "gorm.io/gorm"

type Recommendation struct {
	gorm.Model
	Tag        string   `gorm:"size:100;not null"`
	ComputerID *uint    `gorm:"not null"`
	Computer   Computer `gorm:"foreignKey:ComputerID"`
	Comment    string   `gorm:"size:400;not null"`
}
