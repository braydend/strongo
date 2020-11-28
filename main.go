package main

import (
	"fmt"
	"os"
	"strongo/utils"
	"strongo/utils/firebase"
	"strongo/utils/router"

	fb "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var firebaseApp *fb.App
	var authClient *auth.Client
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
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		panic("Error loading environment variables")
	}

	if utils.IsFixturesEnabled(args) {
		utils.AddFixtureData(db)
	}

	firebaseApp = firebase.GetApp()
	authClient = firebase.GetAuthClient(firebaseApp)

	startServer(db, authClient)

}

// startServer - Boot the API server and listen on port 8080
func startServer(db *gorm.DB, auth *auth.Client) {
	r := gin.Default()
	router.LoadRouter(r, db, auth)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
