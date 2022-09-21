package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test-API/test-API/models"
)

type CreatePostInput struct {
	UserId int    `json:"userId" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

type UpdatePostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// displaying all the posts in the database
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	models.Database.Find(&posts)
	c.IndentedJSON(http.StatusOK, posts)

}

// Creating a post and storing it in the database
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

// Getting all the posts by a specific user ID
func GetPostsByUserId(c *gin.Context) {
	var posts []models.Post
	id, convError := strconv.Atoi(c.Param("userId"))

	if convError == nil {
		models.Database.Where("user_Id = ?", id).Find(&posts)
		if len(posts) == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No posts from this user"})
			return
		}
		c.IndentedJSON(http.StatusOK, posts)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid user id"})
}

// Get specific post according to its ID by a specific user ID
func GetPostsByPostId(c *gin.Context) {
	var post models.Post

	uId, uConvError := strconv.Atoi(c.Param("userId"))
	pId, pConvError := strconv.Atoi(c.Param("id"))

	if uConvError != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Format of user id is invalid"})
	}

	if pConvError != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Format of post id is invalid"})
	}

	if err := models.Database.Where("user_Id = ?", uId).Where("id = ?", pId).Find(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)

}

// Update the title and body of the post
func UpdatePost(c *gin.Context) {
	var post models.Post

	uId, uConvError := strconv.Atoi(c.Param("userId"))
	pId, pConvError := strconv.Atoi(c.Param("id"))

	if uConvError != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Format of user id is invalid"})
	}

	if pConvError != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Format of post id is invalid"})
	}

	if err := models.Database.Where("user_Id = ?", uId).Where("id = ?", pId).Find(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}

	// Validate input
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Database.Model(&post).Updates(input)

	c.IndentedJSON(http.StatusOK, post)

}

// Delete posts
