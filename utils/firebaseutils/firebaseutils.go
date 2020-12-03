package firebaseutils

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

// GetFirebaseAppInstance - Will initiate a connection to Firebase and return a Firebase App instance
func GetFirebaseAppInstance() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	return app
}
