package router

import (
	"devcord-app/controller"
	"devcord-app/repository"
	"devcord-app/firebase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	postRepo := &repository.FirestorePostRepository{
		Client: firebase.Client,
	}

	postController := &controller.PostController{
		Repo: postRepo,
	}

	r.GET("/posts", postController.GetPosts)
	r.POST("/posts", postController.CreatePost)

	return r
}