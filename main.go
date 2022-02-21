package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
	"gitlab.ci.emalify.com/roamtech/asset_be/routes"
)

func main() {
	//load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	api := app.Group("/api/v1")

	//connect to Database
	database.Connect()
	routes.RegisterRoutes(api)

	log.Fatal(app.Listen(":8000"))

}
