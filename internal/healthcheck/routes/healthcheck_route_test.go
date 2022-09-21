package routes_test

import (
	"fmt"
	"net/http/httptest"
	"template-service-go/internal/configuration"
	"testing"

	healthcheckRoutes "template-service-go/internal/healthcheck/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAliveRoute(t *testing.T) {
	config, _ := configuration.Load()
	app := fiber.New()

	healthcheckRoutes.RegisterRoutes(app, config)

	t.Run("It should return status 200", func(t *testing.T) {
		resp, err := app.Test(httptest.NewRequest("GET", "/alive", nil))
		assert.Nil(t, err, fmt.Sprintf("Expected no error, but got '%s'", err))
		assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Expected status code 200, but got %d", resp.StatusCode)
	})
}
