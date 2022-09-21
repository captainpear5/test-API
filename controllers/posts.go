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
		if err := models.Database.Where("user_Id = ?", id).Find(&posts).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "This user has no posts"})
			return
		}
		c.IndentedJSON(http.StatusOK, posts)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid user id"})
}

// Get specific post according to its ID by a specific user ID
func GetPostsByPostId(c *gin.Context) {
	var posts []models.Post

	uId, uConvError := strconv.Atoi(c.Param("userId"))
	pId, pConvError := strconv.Atoi(c.Param("id"))

	if uConvError == nil {
		if err := models.Database.Where("user_Id = ?", uId).Find(&posts).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "This user has no posts"})
			return
		}
		if pConvError == nil {
			for i, p := range posts {
				if p.Id == pId {
					c.IndentedJSON(http.StatusOK, posts[i])
					return
				}
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid post id"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid user id"})
}

// Update the title and body of the post

// Delete posts
