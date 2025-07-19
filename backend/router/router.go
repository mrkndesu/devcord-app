package router

import (
	"github.com/mrkndesu/devcord-app/backend/controller"
	"github.com/mrkndesu/devcord-app/backend/repository"
	"github.com/mrkndesu/devcord-app/backend/firebase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Ginのデフォルトルーターを作成

	// Firestoreクライアントを使ったリポジトリを生成
	postRepo := &repository.FirestorePostRepository{
		Client: firebase.Client,
	}

	// コントローラーにリポジトリをセット
	postController := &controller.PostController{
		Repo: postRepo,
	}

	// ルーティング設定
	r.GET("/posts", postController.GetPosts)             // 投稿一覧取得のエンドポイント
	r.POST("/posts", postController.CreatePost)          // 投稿作成のエンドポイント
	r.GET("/posts/:id", postController.GetPostByID)      // ID指定で投稿取得のエンドポイント
	r.PUT("/posts/:id", postController.UpdatePost)       // ID指定で投稿更新のエンドポイント
	r.DELETE("/posts/:id", postController.DeletePost)    // ID指定で投稿削除のエンドポイント
	r.DELETE("/posts", postController.DeleteAllPosts)    // 全投稿一括削除のエンドポイント

	return r // ルーターを返す
}
