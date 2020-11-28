package router

import (
	"net/http"
	"strongo/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoadRouter - Configure the routes for the webserver
func LoadRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/exercises", (controllers.HandleExercises(db)))
	router.POST("/exercise/create", (controllers.HandleCreateExercise(db)))
	router.POST("/exercises/:exerciseId/createSet", (controllers.HandleCreateSet(db)))
	router.POST("/sets/:setId/update", (controllers.HandleUpdateSet(db)))
	router.DELETE("/sets/:setId", controllers.HandleDeleteSet(db))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid route",
		})
	})
}
