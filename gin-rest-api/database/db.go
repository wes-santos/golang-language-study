package database

import (
	"log"

	"github.com/wes-santos/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	dsn := "host=localhost user=root password=root dbname=PG_GIN_REST_API port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate(&models.Student{})
}
