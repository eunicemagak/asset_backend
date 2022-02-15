package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/routes"
)

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	routes.RegisterAssetsRoutes(api)
	log.Fatal(app.Listen(":8000"))

}
