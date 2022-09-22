package database_test

import (
	"os"
	"template-service-go/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupConnection(t *testing.T) {
	t.Run("Connect", func(t *testing.T) {
		t.Run("it should return a DB connection", func(t *testing.T) {
			_ = os.Setenv("DB_HOST", "localhost")
			_ = os.Setenv("DB_USER", "postgres")
			_ = os.Setenv("DB_PASS", "postgres")
			_ = os.Setenv("DB_NAME", "template_service_go")
			_ = os.Setenv("DB_PORT", "5434")

			db, err := database.SetupConnection()
			assert.Nil(t, err)
			assert.NotNil(t, db)
		})
	})
}
