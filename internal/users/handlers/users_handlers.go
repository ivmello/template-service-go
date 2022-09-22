package handlers

import (
	"template-service-go/internal/configuration"
	"template-service-go/internal/users/models"

	"github.com/gofiber/fiber/v2"
)

func GetAll(config *configuration.AppConfig) fiber.Map {
	users := []models.UserModel{}

	return fiber.Map{
		"status": "success",
		"data":   users,
	}
}
