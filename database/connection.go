package database

import (
	//"fmt"
	//"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	/*host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)*/

	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=yiro password=yiro dbname=tech_hub_test port=6666"), &gorm.Config{})

	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
}
