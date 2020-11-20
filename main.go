package main

import (
	"os"
	"strongo/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var err error
	args := os.Args

	if utils.UseTemporaryDB(args) {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("db/dev.db"), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect database")
	}

	if utils.IsMigrationEnabled(args) {
		utils.MigrateDB(db)
	}

	if utils.IsFixturesEnabled(args) {
		utils.AddFixtureData(db)
	}
}
