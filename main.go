package main

import (
	"github.com/Yiro13/tech-hub-backend/database"
	"github.com/Yiro13/tech-hub-backend/routes"
)

func main() {

	database.DBConnection()
	database.Migrate(database.DB)

	r := routes.SetUpRouter()
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
