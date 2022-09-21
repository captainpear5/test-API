package main

import (
	"github.com/gin-gonic/gin"
	"github.com/test-API/test-API/controllers"
	"github.com/test-API/test-API/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/posts", controllers.GetAllPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:userId", controllers.GetPostsByUserId)
	r.GET("/posts/:userId/:id", controllers.GetPostsByPostId)
	r.PATCH("/posts/:userId/:id", controllers.GetPostsByPostId)
	r.Run()
}
