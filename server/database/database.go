package database

import (
	"fiber-project/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// p := config.Config("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 12)

	// if err != nil {
	// 	log.Println("Idiot")
	// }
	// postgres://ickkbirk:21E9apnDKpxlJ1gmTely7c_c6UhSzK44@satao.db.elephantsql.com/ickkbirk
	// dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	dns := "host=satao.db.elephantsql.com user=ickkbirk password=21E9apnDKpxlJ1gmTely7c_c6UhSzK44 dbname=ickkbirk port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&model.Clientinfo{})
	fmt.Println("Connection Opened to Database")
}