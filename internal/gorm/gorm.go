package db

import (
	gmodels "github.com/yosa12978/ElonScrewedUp/internal/gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(conn string) {
	database, err := gorm.Open(sqlite.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&gmodels.TweetGorm{})
	db = database
}

func GetDB() *gorm.DB {
	return db
}
