package main

import (
	"log"
	"os"
	"strongo/utils"
	"strongo/utils/firebaseutils"
	"strongo/utils/router"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	firebase := firebaseutils.GetFirebaseAppInstance()

	if utils.IsFixturesEnabled(args) {
		utils.AddFixtureData(db)
	}

	startServer(db, firebase)
}

// startServer - Boot the API server and listen on port 8080
func startServer(db *gorm.DB, fb *firebase.App) {
	r := gin.Default()
	router.AllowCORS(r, []string{"*"})
	router.JWTFirewall(r, fb)
	router.LoadRouter(r, db, fb)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
