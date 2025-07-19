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

// Client は Firestore クライアントをグローバルに保持
var Client *firestore.Client

// Init は Firebase と Firestore クライアントの初期化を行う
func Init() error {

	// .env ファイルを読み込み環境変数を設定
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	// サービスアカウントキーのパスを環境変数から取得し認証オプションを作成
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	// Firebase アプリの初期化
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("error initializing firebase app: %v\n", err)
		return err
	}

	// Firestore クライアントの生成
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("error initializing firestore: %v\n", err)
		return err
	}

	// グローバル変数に Firestore クライアントをセット
	Client = firestoreClient

	log.Println("🔥 Firestore initialized successfully")
	return nil
}