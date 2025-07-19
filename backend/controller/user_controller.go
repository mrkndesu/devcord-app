package controller

import (
	"net/http"

	"github.com/mrkndesu/devcord-app/backend/model"
	"github.com/mrkndesu/devcord-app/backend/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// UserController はユーザー情報に関するHTTPリクエストを処理するコントローラー
type UserController struct {
	Repo repository.UserRepository // ユーザーデータ操作用リポジトリ
}

// GetUser は指定ユーザーの情報を取得して返す
func (uc *UserController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("userID")

	user, err := uc.Repo.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser はリクエストからユーザー情報を受け取り新規作成する
func (uc *UserController) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := uc.Repo.Create(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers は全ユーザーを取得して返す（管理者用）
func (uc *UserController) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := uc.Repo.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser は指定ユーザーの情報を更新する
func (uc *UserController) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("userID")

	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if id != updatedUser.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	if err := uc.Repo.Update(ctx, id, updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser は指定ユーザーの投稿を全削除し、ユーザーを削除する
func (uc *UserController) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("userID")

	// ユーザーの投稿を全件削除
	postsIter := uc.Repo.(*repository.FirestoreUserRepository).Client.Collection("users").Doc(userID).Collection("posts").Documents(ctx)
	for {
		doc, err := postsIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list posts"})
			return
		}

		if _, err := doc.Ref.Delete(ctx); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
			return
		}
	}

	// 投稿削除後にユーザー削除
	if err := uc.Repo.Delete(ctx, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User and all posts deleted"})
}
