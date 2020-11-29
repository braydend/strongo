package controllers

import (
	"strongo/models"
	"strongo/utils/httputils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleExercises | GET
//
// /exercises - Get all exercises
//
// /exercises?userId=1234 - Get all exercises with sets for user 1234
//
// /exercises?id=123&id=456 - Get exercises with matching ids
func HandleExercises(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var exercises []models.Exercise
		ids, hasIds := c.GetQueryArray("id")
		userID, hasUserID := c.GetQuery("userId")

		if hasIds {
			db.Find(&exercises, ids)
		} else {
			db.Find(&exercises)
		}

		var exerciseJSON []gin.H

		for _, exercise := range exercises {
			if hasUserID {
				exerciseJSON = append(exerciseJSON, gin.H{
					"id":   exercise.ID,
					"name": exercise.Name,
					"sets": exercise.GetSetsForUser(db, userID),
				})
			} else {
				exerciseJSON = append(exerciseJSON, gin.H{
					"id":   exercise.ID,
					"name": exercise.Name,
				})
			}

		}

		httputils.HandleErrorOrSuccessResponse(c, nil, exerciseJSON, nil)
	}
}

// HandleCreateExercise | POST
//
// /exercise/create - Create a new exercise using data in request body
func HandleCreateExercise(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		exercise := models.Exercise{Name: c.PostForm("name")}
		db.Create(&exercise)

		httputils.HandleErrorOrSuccessResponse(c, nil, exercise, nil)
	}
}

//HandleUpdateExercise | POST
//
// Update an exercise
func HandleUpdateExercise(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		exerciseID, e := httputils.GetIntQueryParamValue(c, "exerciseId")
		var exercise models.Exercise
		db.Find(&exercise).Where("ID = ?", exerciseID)

		name := c.DefaultPostForm("name", exercise.Name)

		exercise.SetName(name)

		httputils.HandleErrorOrSuccessResponse(c, e, exercise, func() { db.Save(&exercise) })
	}
}

// HandleDeleteExercise | DELETE
//
// Soft-delete an exercise by id
func HandleDeleteExercise(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		exerciseID, e := httputils.GetIntQueryParamValue(c, "exerciseId")
		var exercise models.Exercise
		db.Find(&exercise).Where("ID = ?", exerciseID)

		httputils.HandleErrorOrSuccessResponse(c, e, exercise, func() { db.Delete(&exercise) })
	}
}
