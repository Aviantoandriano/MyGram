package helpers

import (
	"final-project/structs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (db *gorm.DB, err error) {
	dsn := "postgres://postgres:94730881@localhost:5432/MyGram?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect to db")
	}
	DB = db
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("cannot connect to db")
	}
	sqlDb.SetConnMaxIdleTime(10)
	sqlDb.SetConnMaxLifetime(100)

	db.AutoMigrate(&structs.User{})
	db.AutoMigrate(&structs.Photo{})
	db.AutoMigrate(&structs.Comment{})
	db.AutoMigrate(&structs.SocialMedia{})
	return
}
