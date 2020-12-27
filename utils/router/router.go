package router

import (
	"fmt"
	"log"
	"net/http"
	"strongo/controllers"
	"strongo/utils/httputils"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// JWTFirewall - Firewall all requests behind a valid JWT provided by Firebase Auth
func JWTFirewall(router *gin.Engine, app *firebase.App) {
	router.Use(func(c *gin.Context) {
		client, err := app.Auth(c)
		if err != nil {
			log.Printf("error getting Auth client: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			c.Abort()
		}

		token := httputils.GetAndAuthoriseTokenForRequest(c, client)
		fmt.Printf("User ID: %s\n", token.UID)
	})
}

// AllowCORS - Allow CORS from the allowed origins passed as an argument
func AllowCORS(router *gin.Engine, allowedOrigins []string) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

// LoadRouter - Configure the routes for the webserver
func LoadRouter(router *gin.Engine, db *gorm.DB, fb *firebase.App) {
	router.GET("/exercises", (controllers.HandleExercises(db)))
	router.GET("/exercise/:exerciseId/sets", (controllers.HandleGetSetsForExerciseByUser(db, fb)))
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
