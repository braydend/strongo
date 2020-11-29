package router

import (
	"net/http"
	"strongo/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AllowCORS - Allow CORS from the allowed origins passed as an argument
func AllowCORS(router *gin.Engine, allowedOrigins []string) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

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
