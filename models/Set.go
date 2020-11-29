package models

import (
	"gorm.io/gorm"
)

// Set is a collection of repetitions of an exercise
type Set struct {
	gorm.Model
	Reps       int
	ExerciseID int
	UserID     string
	Weight     float64
	Unit       string
}

func (set *Set) SetReps(reps int) *Set {
	set.Reps = reps
	return set
}

func (set *Set) SetExercise(id int) *Set {
	set.ExerciseID = id
	return set
}

func (set *Set) SetUser(id string) *Set {
	set.UserID = id
	return set
}

func (set *Set) SetWeight(weight float64) *Set {
	set.Weight = weight
	return set
}

func (set *Set) SetUnit(unit string) *Set {
	set.Unit = unit
	return set
}
