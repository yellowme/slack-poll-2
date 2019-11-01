package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var GormDatabase *gorm.DB

func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "poll.db")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&PollModel{})
	db.AutoMigrate(&PollAnswerModel{})

	GormDatabase = db

	return GormDatabase
}
