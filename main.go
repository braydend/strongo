package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strongo/controllers"
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
	port := 8080
	fmt.Println(fmt.Sprintf("Starting server on port: %d", port))

	http.HandleFunc("/exercises", controllers.HandleExercises(db))
	http.HandleFunc("/exercise/create", controllers.HandleCreateExercise(db))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}
