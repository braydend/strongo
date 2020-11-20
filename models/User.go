package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Sets    []Set
	Weights []userWeight
}

type userWeight struct {
	Weight float32
	Unit   string
	UserID uint
}
