package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=dpg-cl8i7olb7ptc73d9gnp0-a user=yiro password=TJbUmAK78EhUHrq9QCIZCcxclZAo9zi9 dbname=tech_hub_n8ut port=5432"), &gorm.Config{})

	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
}
