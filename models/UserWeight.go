package models

import "gorm.io/gorm"

// UserWeight - Data type for storing the user's weight over time
type UserWeight struct {
	gorm.Model
	Weight float64
	Unit   string
	UserID string
}

// SetWeight - Setter for weight
func (u *UserWeight) SetWeight(weight float64) *UserWeight {
	u.Weight = weight
	return u
}

// SetUser - Setter for user
func (u *UserWeight) SetUser(id string) *UserWeight {
	u.UserID = id
	return u
}

// SetUnit - Setter for weight unit
func (u *UserWeight) SetUnit(unit string) *UserWeight {
	u.Unit = unit
	return u
}
