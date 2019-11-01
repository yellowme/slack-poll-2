package test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PollModel struct {
	ID       string
	Question string
	Owner    string
	Options  string
	Mode     string
}

type PollAnswerModel struct {
	ID     string
	Option string
	Owner  string
	Poll   PollModel
	PollID string
}

var testDatabase *gorm.DB

func InitializeTestDatabase() *gorm.DB {
	if testDatabase != nil {
		return testDatabase
	}

	db, err := gorm.Open("sqlite3", "poll_test.db")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&PollModel{})
	db.AutoMigrate(&PollAnswerModel{})

	testDatabase = db

	return db
}
