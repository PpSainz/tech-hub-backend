package models

import "gorm.io/gorm"

type MotherBoard struct {
	gorm.Model
	Brand          string `gorm:"not null;size:50"`
	ModelName      string `gorm:"not null;size:100"`
	FormFactor     string `gorm:"not null;size:100"`
	SocketType     string `gorm:"not null;size:50"`
	RAMSlots       int    `gorm:"not null"`
	MaxMemoryGB    int    `gorm:"not null"`
	PCIeSlots      int    `gorm:"not null"`
	USBPorts       int    `gorm:"not null"`
	SATAConnectors int
	Networking     string `gorm:"size:200"`
	ImageURL       string `gorm:"not null;size:200"`
}
