package main

import (
	"log"
	"template-service-go/internal/configuration"
	healthcheckRoutes "template-service-go/internal/healthcheck/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api := app.Group("/api")

	healthcheckRoutes.RegisterRoutes(api, config)

	log.Fatal(app.Listen(":" + config.Port))
}
