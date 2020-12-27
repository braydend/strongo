package controllers

import (
	"strongo/models"
	"strongo/utils/httputils"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleCreateSet | POST
//
// Create a set for an exercise
func HandleCreateSet(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		exerciseID, e := httputils.GetUintQueryParamValue(c, "exerciseId")
		reps, e := httputils.GetIntPostValue(c, "reps")
		weight, e := httputils.GetFloatPostValue(c, "weight")
		userID := c.PostForm("userId")
		unit := c.PostForm("unit")
		set := models.Set{
			ExerciseID: exerciseID,
			Reps:       reps,
			Weight:     weight,
			UserID:     userID,
			Unit:       unit,
		}

		httputils.HandleErrorOrSuccessResponse(c, e, set, func() { db.Create(&set) })
	}
}

// HandleUpdateSet | POST
//
// Update a set's fields
func HandleUpdateSet(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		setID, e := httputils.GetIntPostValue(c, "setId")
		var set models.Set
		db.Find(&set).Where("ID = ?", setID)

		reps, e := httputils.GetDefaultIntPostValue(c, "reps", set.Reps)
		weight, e := httputils.GetDefaultFloatPostValue(c, "weight", set.Weight)
		userID := c.DefaultPostForm("userId", set.UserID)
		unit := c.DefaultPostForm("unit", set.Unit)

		set.SetReps(reps).SetWeight(weight).SetUser(userID).SetUnit(unit)

		httputils.HandleErrorOrSuccessResponse(c, e, set, func() { db.Save(&set) })
	}
}

// HandleDeleteSet | DELETE
//
// Soft-delete a set by id
func HandleDeleteSet(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		setID, e := httputils.GetIntQueryParamValue(c, "setId")
		var set models.Set
		db.Find(&set).Where("ID = ?", setID)

		httputils.HandleErrorOrSuccessResponse(c, e, set, func() { db.Delete(&set) })
	}
}

// HandleGetSetsForExerciseByUser | GET
//
// Fetch all sets for an exercise using JWT for auth
func HandleGetSetsForExerciseByUser(db *gorm.DB, fb *firebase.App) func(*gin.Context) {
	return func(c *gin.Context) {
		exerciseID := c.Param("exerciseId")
		auth, err := fb.Auth(c)

		token := httputils.GetAuthTokenForRequest(c, auth)

		var exercise models.Exercise
		db.Find(&exercise, exerciseID)
		sets := exercise.GetSetsForUser(db, token.UID)

		httputils.HandleErrorOrSuccessResponse(c, err, sets, nil)
	}
}
