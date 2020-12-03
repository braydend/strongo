package models

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Name string `gorm:"index:idx_name,unique"`
	Sets []Set
}

func (exercise *Exercise) SetName(name string) *Exercise {
	exercise.Name = name
	return exercise
}

func (exercise *Exercise) GetSetsForUser(db *gorm.DB, id string) *[]Set {
	var sets []Set
	db.Model(&exercise).Where("user_id = ?", id).Association("Sets").Find(&sets)

	return &sets
}
