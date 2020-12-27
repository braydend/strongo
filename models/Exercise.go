package models

import (
	"gorm.io/gorm"
)

// Exercise - Model for working with exercises
type Exercise struct {
	gorm.Model
	Name string `gorm:"index:idx_name,unique"`
	Sets []Set
}

// SetName - Setter for name field on an Exercise
func (exercise *Exercise) SetName(name string) *Exercise {
	exercise.Name = name
	return exercise
}

// GetSetsForUser - Getter for Sets created by the given User on the Exercise
func (exercise *Exercise) GetSetsForUser(db *gorm.DB, userID string) *[]Set {
	var sets []Set
	db.Model(&exercise).Where("user_id = ?", userID).Association("Sets").Find(&sets)

	return &sets
}
