package controller

import (
	"net/http"
	"time"

	"github.com/mrkndesu/devcord-app/backend/model"
	"github.com/mrkndesu/devcord-app/backend/repository"

	"github.com/gin-gonic/gin"
)

// PostController は投稿に関するHTTPリクエストを処理するコントローラー
type PostController struct {
	Repo repository.PostRepository // 投稿データ操作用リポジトリ
}

// GetPosts は指定ユーザーの全投稿を取得して返す
func (pc *PostController) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	posts, err := pc.Repo.GetAll(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// CreatePost は指定ユーザーに新しい投稿を作成する
func (pc *PostController) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID is required in path"})
		return
	}

	post.UserID = userID
	post.CreatedAt = time.Now()

	if err := pc.Repo.Create(ctx, userID, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPostByID は指定ユーザーの特定投稿を取得して返す
func (pc *PostController) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()
	postID := c.Param("postID")
	userID := c.Param("userID")

	if userID == "" || postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and post_id are required"})
		return
	}

	post, err := pc.Repo.GetByID(ctx, userID, postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost は指定ユーザーの特定投稿を更新する
func (pc *PostController) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := c.Param("userID")
	postID := c.Param("postID")

	if userID == "" || postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID and postID are required"})
		return
	}

	post.UserID = userID
	post.ID = postID

	if err := pc.Repo.Update(ctx, userID, postID, post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost は指定ユーザーの特定投稿を削除する
func (pc *PostController) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()
	postID := c.Param("postID")
	userID := c.Param("userID")

	if userID == "" || postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and post_id are required"})
		return
	}

	post, err := pc.Repo.GetByID(ctx, userID, postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := pc.Repo.Delete(ctx, userID, post.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// DeleteAllPosts は指定ユーザーの全投稿を一括削除する
func (pc *PostController) DeleteAllPosts(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if err := pc.Repo.DeleteAll(ctx, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All posts deleted"})
}
