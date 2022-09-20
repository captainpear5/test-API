package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/test-API/test-API/models"
)

type CreatePostInput struct {
	UserId int    `json:"UserId" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

// displaying all the posts in the database
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	models.Database.Find(&posts)
	c.IndentedJSON(http.StatusOK, posts)

}

func CreatePost(c *gin.Context) {
	var input CreatePostInput

	// whats the difference between BindJSON and ShouldBind
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{UserId: input.UserId, Title: input.Title, Body: input.Body}
	models.Database.Create(&post)
	c.IndentedJSON(http.StatusCreated, input)

}
