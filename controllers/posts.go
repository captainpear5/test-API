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
func getPostsByUserId(c *gin.Context) {
	var post models.Post

	if err := models.Database.Where("UserId = ?", c.Param("userId")).Find(&post).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such user exists"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}

/*
func getPostsByUserId(id int) (*gorm.DB, error) {
	var post models.Post

	userPosts := models.Database.Where("UserId = ?", id).Find(&post)

	if userPosts != nil {
		return userPosts, nil
	}

	return nil, errors.New("no posts with this user id")
}

func readPostsByUserId(c *gin.Context) {
	id, convError := strconv.Atoi(c.Param("UserId"))


}
*/
