package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// GetApp - Set the firebase App instance
func GetApp() *firebase.App {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_PATH"))
	config := &firebase.Config{ProjectID: "strongo-b5024"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

// GetAuthClient - Get a firebase auth client instance
func GetAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}

// GetUserByUID - Find a user with their UID
func GetUserByUID(client *auth.Client, uid string) (*auth.UserRecord, error) {
	u, err := client.GetUser(context.Background(), "test@email.co")
	return u, err
}
