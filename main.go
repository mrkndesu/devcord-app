package main

import (
	"log"
	"net/http"

	"github.com/mrkndesu/devcord-app/backend/firebase"
	"github.com/mrkndesu/devcord-app/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Firebase初期化、失敗したらログ出力して終了
	if err := firebase.Init(); err != nil {
		log.Fatalf("🔥 Firebase init failed: %v", err)
	}

	// ルーターのセットアップ
	r := router.SetupRouter()

	// ルートパスにアクセスされたらWelcomeメッセージを返す
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Devcord API!",
		})
	})

	// サーバー起動、失敗したらログ出力して終了
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("🚫 Server failed to start: %v", err)
	}
}