package database

import (
	"fmt"
	"go-fiber-jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "satao.db.elephantsql.com"
	port     = 5432
	user     = "komskzvi"
	password = "tEvrs9-e8Rk3hrpFQlxeo9Kd3drXmLsh" //Enter your password for the DB
	dbname   = "komskzvi"
)

// "host=satao.db.elephantsql.com user=komskzvi password=tEvrs9-e8Rk3hrpFQlxeo9Kd3drXmLsh dbname=komskzvi port=5432 sslmode=disable"
// postgres://komskzvi:tEvrs9-e8Rk3hrpFQlxeo9Kd3drXmLsh@satao.db.elephantsql.com/komskzvi
var dsn string = fmt.Sprintf("host=%s user=%s   password=%s dbname=%s port=%d sslmode=disable ",
	host, user, password, dbname, port)

var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	DB.AutoMigrate(&models.User{}) // we are going to create a models.go file for the User Model.
}
