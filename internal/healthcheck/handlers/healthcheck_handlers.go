package handlers

import (
	"template-service-go/internal/configuration"

	"github.com/gofiber/fiber/v2"
)

func Alive(config *configuration.AppConfig) fiber.Map {
	return fiber.Map{
		"status": "success",
		"data":   "alive",
	}
}
