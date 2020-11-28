package main

import (
	"os"
	"strongo/utils"
	"strongo/utils/router"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var err error
	args := os.Args

	if utils.UseTemporaryDB(args) {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

		// Always run migration when using temporary DB
		utils.MigrateDB(db)
	} else {
		db, err = gorm.Open(sqlite.Open("db/dev.db"), &gorm.Config{})

		if utils.IsMigrationEnabled(args) {
			utils.MigrateDB(db)
		}
	}

	if err != nil {
		panic("failed to connect database")
	}

	if utils.IsFixturesEnabled(args) {
		utils.AddFixtureData(db)
	}

	startServer(db)
}

// startServer - Boot the API server and listen on port 8080
func startServer(db *gorm.DB) {
	r := gin.Default()
	router.LoadRouter(r, db)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
