package handlers_test

import (
	"template-service-go/internal/configuration"
	usersHandlers "template-service-go/internal/users/handlers"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	config, _ := configuration.Load()

	t.Run("GetAll", func(t *testing.T) {
		t.Run("it should return a list of users", func(t *testing.T) {

			received := usersHandlers.GetAll(config)
			expected := fiber.Map{
				"status": "success",
				"data": map[string]interface{}{
					"id":    "00000000000000000000000000000000",
					"name":  "teste",
					"email": "teste@example.com",
				},
			}
			assert.Equal(t, expected, received)
		})
	})
}
