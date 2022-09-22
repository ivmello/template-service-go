package main

import (
	"log"
	"template-service-go/internal/configuration"
	healthcheckRoutes "template-service-go/internal/healthcheck/routes"
	usersRoutes "template-service-go/internal/users/routes"
	"template-service-go/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api := app.Group("/api")

	_, err = database.SetupConnection()
	if err != nil {
		log.Fatal(err)
	}

	healthcheckRoutes.RegisterRoutes(api, config)
	usersRoutes.RegisterRoutes(api, config)

	log.Fatal(app.Listen(":" + config.Port))
}
