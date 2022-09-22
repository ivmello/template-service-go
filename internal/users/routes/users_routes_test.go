package routes_test

import (
	"fmt"
	"net/http/httptest"
	"template-service-go/internal/configuration"
	"testing"

	usersRoutes "template-service-go/internal/users/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllRoute(t *testing.T) {
	config, _ := configuration.Load()
	app := fiber.New()

	usersRoutes.RegisterRoutes(app, config)

	t.Run("it should return status OK (200)", func(t *testing.T) {
		response, err := app.Test(httptest.NewRequest("GET", "/users", nil))
		assert.Nil(t, err, fmt.Sprintf("Expected no error, but got '%s'", err))
		assert.Equal(t, fiber.StatusOK, response.StatusCode, "Error: %d", response.StatusCode)
	})
}
