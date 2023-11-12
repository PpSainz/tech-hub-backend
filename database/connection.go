package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=yiro password=yiro dbname=tech_hub_test port=3333"), &gorm.Config{})

	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
}
