package configuration_test

import (
	"fmt"
	"os"
	"template-service-go/internal/configuration"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	t.Run("returns an env string", func(t *testing.T) {
		_ = os.Setenv("TEST_ENV_STRING", "teste")
		value := configuration.GetEnvString("TEST_ENV_STRING")
		assert.Equal(t, "teste", value, fmt.Sprintf("Expected value to be 'test' but got '%s'", value))
	})
	defer func() {
		_ = os.Unsetenv("TEST_ENV_STRING")
	}()
}

func TestGetEnvBool(t *testing.T) {
	t.Run("returns an env boolean", func(t *testing.T) {
		_ = os.Setenv("TEST_ENV_BOOL", "true")
		value := configuration.GetEnvBool("TEST_ENV_BOOL")
		assert.Equal(t, true, value, fmt.Sprintf("Expected value to be true but got '%t'", value))
	})
	defer func() {
		_ = os.Unsetenv("TEST_ENV_BOOL")
	}()
}

func TestLoad(t *testing.T) {
	t.Run("returns an error if the file does not exist", func(t *testing.T) {
		_ = os.Setenv("APP_PORT", "3000")
		_ = os.Setenv("APP_ENV", "development")
		_ = os.Setenv("APP_TESTING", "true")

		config, err := configuration.Load()

		t.Run("it should return load configurations without error", func(t *testing.T) {
			assert.Nil(t, err, fmt.Sprintf("Expected no error, but got %s", err))
		})

		t.Run("it should return configuration struct", func(t *testing.T) {
			assert.Equal(t, "3000", config.Port)
			assert.Equal(t, true, config.IsTesting, fmt.Sprintf("Expected value to be 'true', but got %t", config.IsTesting))
		})

		t.Run("it should return an error when there's a missing required variable", func(t *testing.T) {
			_ = os.Unsetenv("APP_PORT")
			_, err := configuration.Load()
			assert.NotNil(t, err, "Expected error, but got nil")
		})
	})

	defer func() {
		_ = os.Unsetenv("APP_PORT")
		_ = os.Unsetenv("APP_ENV")
		_ = os.Unsetenv("APP_TESTING")
	}()
}
