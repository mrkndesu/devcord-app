package controller

import (
	"net/http"

	"devcord-app/model"
	"devcord-app/repository"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Repo repository.PostRepository
}

// GET /posts
func (pc *PostController) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()

	posts, err := pc.Repo.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// POST /posts
func (pc *PostController) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := pc.Repo.Create(ctx, post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post created"})
}