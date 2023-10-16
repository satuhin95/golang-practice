package initial

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// postgres://bxutzhek:d5q-48EfHfcf2i2YaQe-gpgB0gog5RYG@satao.db.elephantsql.com/bxutzhek
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}
}
