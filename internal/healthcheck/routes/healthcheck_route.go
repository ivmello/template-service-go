package routes

import (
	"template-service-go/internal/configuration"

	healthcheckHandler "template-service-go/internal/healthcheck/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router, config *configuration.AppConfig) {
	app.Get("/alive", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(healthcheckHandler.Alive(config))
	})
}
