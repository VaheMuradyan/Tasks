package controllers

import (
	"go-redis/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAuthorAndBook(c *gin.Context) {

	var body struct {
		Author string
		Book   string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query body is incorrect"})
	}

	err := initializers.RedisClient.Set(initializers.Ctx, body.Author, body.Book, 0).Err()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query body is incorrect"})
	}

	c.Status(http.StatusCreated)
}

func GetBook(c *gin.Context) {
	var body struct {
		Author string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query body is incorrect"})
	}

	book, err := initializers.RedisClient.Get(initializers.Ctx, body.Author).Result()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}

	c.IndentedJSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {

	var body struct {
		Author string
	}

	var books []string
	keys := []string{}

	iter := initializers.RedisClient.Scan(initializers.Ctx, 0, body.Author, 0).Iterator()
	for iter.Next(initializers.Ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "books not found"})
	}

	for _, val := range keys {
		book, err := initializers.RedisClient.Get(initializers.Ctx, val).Result()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "booksnot found"})
		}
		books = append(books, book)
	}

	c.IndentedJSON(http.StatusOK, books)
}

func DeleteBook(c *gin.Context) {
	var body struct {
		Author string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query body is incorrect"})
	}

	result := initializers.RedisClient.Del(initializers.Ctx, body.Author)
	if result.Err() != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book don't deleted"})
	}

	if result.Val() == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book don't deleted"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
