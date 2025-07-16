package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)

var Client *firestore.Client

// 初期化
func Init() error {

	// .env読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("error initializing firebase app: %v\n", err)
		return err
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("error initializing firestore: %v\n", err)
		return err
	}

	Client = firestoreClient
	log.Println("🔥 Firestore initialized successfully")
	return nil

}