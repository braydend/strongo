package models

import (
	"gorm.io/gorm"
)

// Set is a collection of repetitions of an exercise
type Set struct {
	gorm.Model
	Reps       int8
	ExerciseID uint
	UserID     uint
	Weight     float32
	Unit       string
}
