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
