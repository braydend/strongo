package controllers

import (
	"fmt"
	"net/http"
	"strongo/models"
	"strongo/utils"

	"gorm.io/gorm"
)

// HandleExercises - GET -
//
// /exercises - Get all exercises
//
// /exercises?id=123&?id=456 - Get exercises with matching ids
func HandleExercises(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := utils.RestrictMethods([]string{utils.GET}, r.Method)
		if err != nil {
			w.WriteHeader(405)
			fmt.Fprintf(w, err.Error())
			return
		}
		var exercises []models.Exercise
		ids := utils.GetQueryParamValues(r.URL, "id", nil)

		if ids == nil {
			db.Find(&exercises)
		} else {
			db.Find(&exercises, ids)
		}

		for _, exercise := range exercises {
			fmt.Fprintf(w, fmt.Sprintf("%d: %s \n", exercise.ID, exercise.Name))
		}
	}
}

// HandleCreateExercise - POST -
//
// /exercise/create - Create a new exercise using data in request body
func HandleCreateExercise(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := utils.RestrictMethods([]string{utils.POST}, r.Method)
		if err != nil {
			w.WriteHeader(405)
			fmt.Fprintf(w, err.Error())
			return
		}
		r.ParseForm()
		exercise := models.Exercise{Name: r.Form.Get("name")}
		db.Create(&exercise)
		fmt.Fprintf(w, fmt.Sprintf("Exercise %d: %s", exercise.ID, exercise.Name))
	}
}
