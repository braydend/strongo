package firebaseutils

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// GetFirebaseAppInstance - Will initiate a connection to Firebase and return a Firebase App instance
func GetFirebaseAppInstance() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	return app
}

// AuthoriseFirebaseToken - Attempt to authorise a given token against Firebase auth
func AuthoriseFirebaseToken(c *gin.Context, auth *auth.Client, idToken string) *auth.Token {
	token, err := auth.VerifyIDToken(c, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		c.Abort()
	}
	return token
}
