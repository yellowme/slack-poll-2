package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Database *GormDatabase

type GormDatabase struct {
	DB *gorm.DB
}

func InitializeDatabase() *GormDatabase {
	db, err := gorm.Open("sqlite3", "poll.db")

	if err != nil {
		panic("failed to connect database")
	}

	GormDatabase := &GormDatabase{
		DB: db,
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&PollModel{})
	db.AutoMigrate(&PollAnswerModel{})

	return GormDatabase
}
