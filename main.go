package main

import (
	"os"

	models "blog_test/models"
	routers "blog_test/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()          // Load env variables
	models.ConnectDataBase() // load db
	
	router := routers.SetupRouter()

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "9000"
	}

	router.Run(":" + port)
}
