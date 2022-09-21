package handlers_test

import (
	"template-service-go/internal/configuration"
	healthcheckHandler "template-service-go/internal/healthcheck/handlers"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAlive(t *testing.T) {
	config, _ := configuration.Load()

	t.Run("Alive", func(t *testing.T) {
		t.Run("it should return json", func(t *testing.T) {
			received := healthcheckHandler.Alive(config)
			expected := fiber.Map{
				"status": "success",
				"data":   "alive",
			}
			assert.Equal(t, expected, received)
		})
	})
}
