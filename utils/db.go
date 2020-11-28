package utils

import (
	"fmt"
	"strongo/models"

	"gorm.io/gorm"
)

// MigrateDB - Add models to list when they are added
func MigrateDB(db *gorm.DB) bool {
	fmt.Println("---Running Migrations---")
	exerciseErr := db.AutoMigrate(&models.Exercise{})
	setErr := db.AutoMigrate(&models.Set{})
	userErr := db.AutoMigrate(&models.User{})

	if exerciseErr != nil {
		fmt.Println("Exercise error:")
		fmt.Println(exerciseErr.Error())
		return false
	}

	if setErr != nil {
		fmt.Println("Set error:")
		fmt.Println(setErr.Error())
		return false
	}

	if userErr != nil {
		fmt.Println("User error:")
		fmt.Println(userErr.Error())
		return false
	}

	return true
}

func generateFixtureExercises() []models.Exercise {
	exerciseNames := []string{"Bench Press", "Bicep Curl"}
	var exercises []models.Exercise

	for _, name := range exerciseNames {
		exercises = append(exercises, models.Exercise{Name: name})
	}

	return exercises
}

// AddFixtureData - Store fixture data in db
func AddFixtureData(db *gorm.DB) bool {
	fmt.Println("---Running Fixtures---")
	exercises := generateFixtureExercises()

	tx := db.Create(&exercises)

	if tx.Error != nil {
		return false
	}

	return true
}
