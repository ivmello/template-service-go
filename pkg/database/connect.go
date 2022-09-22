package database

import (
	"fmt"
	"template-service-go/internal/users/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupConnection() (*gorm.DB, error) {
	// host := configuration.GetEnvString("DB_HOST")
	// user := configuration.GetEnvString("DB_USER")
	// pass := configuration.GetEnvString("DB_PASS")
	// dbname := configuration.GetEnvString("DB_NAME")
	// port, _ := strconv.Atoi(configuration.GetEnvString("DB_PORT"))

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
	// 	host,
	// 	user,
	// 	pass,
	// 	dbname,
	// 	port,
	// )

	dsn := "host=db user=postgres password=postgres dbname=template_service_go port=5435 sslmode=disable TimeZone=America/Sao_Paulo"

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.UserModel{})

	return db, nil
}
