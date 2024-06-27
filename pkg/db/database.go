package db

import (
	"fmt"
	"payd/internal/auth/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	//host := os.Getenv("POSTGRES_HOST")
	//user := os.Getenv("POSTGRES_USER")
	//password := os.Getenv("POSTGRES_PASSWORD")
	//dbname := os.Getenv("POSTGRES_DB")
	//port := os.Getenv("POSTRES_PORT")

	dsn := "host=127.0.0.1 user=postgres password=1234 dbname=paydtest port=5432 sslmode=disable"
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected successfully")
	DB.AutoMigrate(&model.User{})
}
