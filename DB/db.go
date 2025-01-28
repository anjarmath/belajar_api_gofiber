package db

import (
	"os"

	"github.com/anjarmath/01_golang_api_sederhana/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	conn_url := os.Getenv("DATABASE_URL")

	database, err := gorm.Open(postgres.Open(conn_url), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database")
	}
	database.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	database.AutoMigrate(
		&model.Book{},
		&model.User{},
	)

	DB = database
}
