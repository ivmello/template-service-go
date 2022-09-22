package routes

import (
	"template-service-go/internal/configuration"

	usersHandlers "template-service-go/internal/users/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, config *configuration.AppConfig) {
	routes := router.Group("/users")

	routes.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(usersHandlers.GetAll(config))
	})
}
