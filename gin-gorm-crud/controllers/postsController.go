package controllers

import (
	"net/http"

	"github.com/VaheMuradyan/Tasks/gin-gorm-crud/initializers"
	"github.com/VaheMuradyan/Tasks/gin-gorm-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query body is incorrect"})
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "post was not created"})
	}

	c.IndentedJSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.IndentedJSON(http.StatusOK, posts)
}

func GetPostById(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
	}

	c.IndentedJSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query body is incorrect"})
	}

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
	}

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.IndentedJSON(http.StatusCreated, post)
}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
	}

	c.Status(http.StatusOK)
}
