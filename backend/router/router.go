package router

import (
	"github.com/mrkndesu/devcord-app/backend/controller"
	"github.com/mrkndesu/devcord-app/backend/repository"
	"github.com/mrkndesu/devcord-app/backend/firebase"

	"github.com/gin-gonic/gin"
)

// SetupRouter は Gin のルーターを初期化し、各エンドポイントにコントローラを紐づけて返す
func SetupRouter() *gin.Engine {
	r := gin.Default() // デフォルトのルーター

	// =========================
	// ユーザーリポジトリとコントローラの初期化
	// =========================

	// Firestore クライアントを使ってユーザーリポジトリを初期化
	userRepo := &repository.FirestoreUserRepository{
		Client: firebase.Client,
	}

	// ユーザーコントローラを初期化（リポジトリを注入）
	userController := &controller.UserController{
		Repo: userRepo,
	}

	// =========================
	// 投稿リポジトリとコントローラの初期化
	// =========================

	// Firestore クライアントを使って投稿リポジトリを初期化
	postRepo := &repository.FirestorePostRepository{
		Client: firebase.Client,
	}

	// 投稿コントローラを初期化（リポジトリを注入）
	postController := &controller.PostController{
		Repo: postRepo,
	}

	// =========================
	// ユーザー関連のルーティング
	// =========================

	r.POST("/users", userController.CreateUser)         // ユーザー作成
	r.GET("/users", userController.GetUsers)            // 全ユーザー取得（管理者用）
	r.GET("/users/:userID", userController.GetUser)     // 指定ユーザー取得
	r.PUT("/users/:userID", userController.UpdateUser)  // 指定ユーザー更新
	r.DELETE("/users/:userID", userController.DeleteUser) // 指定ユーザー削除

	// =========================
	// 投稿関連のルーティング
	// =========================

	r.GET("/users/:userID/posts", postController.GetPosts)               // 全投稿取得
	r.POST("/users/:userID/posts", postController.CreatePost)            // 投稿作成
	r.GET("/users/:userID/posts/:postID", postController.GetPostByID)    // 投稿取得
	r.PUT("/users/:userID/posts/:postID", postController.UpdatePost)     // 投稿更新
	r.DELETE("/users/:userID/posts/:postID", postController.DeletePost)  // 投稿削除
	r.DELETE("/users/:userID/posts", postController.DeleteAllPosts)      // 全投稿削除

	return r
}
