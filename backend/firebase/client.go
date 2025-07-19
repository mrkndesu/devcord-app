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

// Client は Firestore クライアントをグローバルに保持する変数
var Client *firestore.Client

// Init は Firebase と Firestore クライアントを初期化する
// .env から環境変数を読み込み、認証情報を使って Firebase アプリと Firestore クライアントを作成し、
// グローバル変数 Client にセットする
// エラーがあればログに出力して返す
func Init() error {
	// .env ファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	// 認証情報ファイルのパスを環境変数から取得
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	// Firebase アプリを初期化
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("error initializing firebase app: %v\n", err)
		return err
	}

	// Firestore クライアントを生成
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("error initializing firestore: %v\n", err)
		return err
	}

	// グローバル変数にセット
	Client = firestoreClient

	log.Println("🔥 Firestore initialized successfully")
	return nil
}
