package main

import (
	"log"
	"net/http"

	"devcord-app/firebase"
	"devcord-app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := firebase.Init(); err != nil {
		log.Fatalf("🔥 Firebase init failed: %v", err)
	}

	r := router.SetupRouter()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Welcome to Devcord API!",
    })
  })

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("🚫 Server failed to start: %v", err)
	}
}