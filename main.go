package main

import (
	"log"
	"net/http"

	"github.com/mrkndesu/devcord-app/backend/firebase"
	"github.com/mrkndesu/devcord-app/backend/router"

	"github.com/gin-gonic/gin"
)

// main はアプリの起動処理
func main() {
	// Firebase を初期化。失敗したら終了
	if err := firebase.Init(); err != nil {
		log.Fatalf("🔥 Firebase init failed: %v", err)
	}

	// ルーターを設定
	r := router.SetupRouter()

	// "/" にアクセスされたときにメッセージを返す
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Devcord API!",
		})
	})

	// ポート8080でサーバーを起動。失敗したら終了
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("🚫 Server failed to start: %v", err)
	}
}

