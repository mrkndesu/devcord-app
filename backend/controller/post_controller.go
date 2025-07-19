package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/mrkndesu/devcord-app/backend/model"
	"github.com/mrkndesu/devcord-app/backend/repository"

	"github.com/gin-gonic/gin"
)

// PostController は投稿に関するHTTPリクエストを処理するコントローラー
type PostController struct {
	Repo repository.PostRepository // 投稿操作用のリポジトリインターフェース
}

// GetPosts は GET /posts のハンドラー
func (pc *PostController) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()

	// 投稿一覧を取得
	posts, err := pc.Repo.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	// 投稿一覧をJSONで返す
	c.JSON(http.StatusOK, posts)
}

// CreatePost は POST /posts のハンドラー
func (pc *PostController) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var post model.Post
	// リクエストのJSONをPost構造体にバインド
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 作成日時を現在時刻にセット
	post.CreatedAt = time.Now()

	// 投稿をFirestoreに保存し、IDをpostにセットする
	if err := pc.Repo.Create(ctx, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// 作成した投稿内容をレスポンスで返す
	c.JSON(http.StatusOK, post)
}

// GET /posts/:id
// 指定されたIDの投稿をFirestoreから取得し、JSONで返す
func (pc *PostController) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id") // URLパラメータからID取得

	post, err := pc.Repo.GetByID(ctx, id) // Firestoreから投稿取得
	if err != nil {
		// 投稿が見つからなければ404返す
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	// 投稿が見つかれば200で返す
	c.JSON(http.StatusOK, post)
}

// PUT /posts/:id
// リクエストボディのJSONをデコードし、指定IDの投稿を更新する
func (pc *PostController) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id") // URLパラメータからID取得

	var updatedPost model.Post
	// JSONのバインドに失敗したら400エラー返す
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Firestore上で投稿を更新
	if err := pc.Repo.Update(ctx, id, updatedPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	// 成功時はJSONでメッセージ返却
	c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
}

// DELETE /posts/:id
// 指定IDの投稿が存在するか確認し、あれば削除する
func (pc *PostController) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id") // URLパラメータからID取得

	// 投稿の存在チェック
	post, err := pc.Repo.GetByID(ctx, id)
	if err != nil {
		if err.Error() == "not found" {
			// 投稿がなければ404返す
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		// その他エラーは500返す
		log.Printf("failed to get post with id %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	// 投稿が存在する場合、削除を実行
	if err := pc.Repo.Delete(ctx, post.ID); err != nil {
		log.Printf("failed to delete post with id %s: %v", post.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	// 削除成功のメッセージ返却
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// DELETE /posts
// 全投稿を一括削除する
func (pc *PostController) DeleteAllPosts(c *gin.Context) {
	ctx := c.Request.Context()

	// Firestore上の全投稿を削除
	if err := pc.Repo.DeleteAll(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all posts"})
		return
	}

	// 削除成功のメッセージ返却
	c.JSON(http.StatusOK, gin.H{"message": "All posts deleted"})
}