package initial

import "github.com/satuhin95/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
