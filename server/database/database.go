package database

import (
	"fmt"
	"react-go-fiber-jwt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  var DBConn *gorm.DB
  func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/react_gofiber_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("DB Connection Error")
	}
	db.AutoMigrate(&models.User{})
	DBConn = db

	fmt.Println("Database Connection Successful")
  }