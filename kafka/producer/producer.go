package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	r := gin.Default()
	r.POST("/comments", createComment)
	r.Run(":3000")
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:29092"}
}

func createComment(c *gin.Context) {
	cmt := new(Comment)
	if err := c.BindJSON(&cmt); err != nil {
		log.Println(err)
		c.Status(400)
	}

	cmtInBytes, err := json.Marshal(cmt)

	if err != nil {
		log.Println(err)
		c.Status(400)
	}

	PushCommentToQueue("comments", cmtInBytes)

	c.IndentedJSON(http.StatusOK, cmt)

}
