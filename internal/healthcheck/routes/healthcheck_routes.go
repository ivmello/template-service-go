package routes

import (
	"template-service-go/internal/configuration"

	healthcheckHandlers "template-service-go/internal/healthcheck/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, config *configuration.AppConfig) {
	router.Get("/alive", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(healthcheckHandlers.Alive(config))
	})
}
