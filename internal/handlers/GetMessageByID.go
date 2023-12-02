package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// jwt token
type MessagePattern struct {
	Content string `json:"content"`
}

type jso struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var jsonic = []jso{
	{ID: "1", Content: "Hello"},
	{ID: "2", Content: "Bye"},
}

func GetMessageByID(c *gin.Context) {
	id := c.GetHeader("id")
	time.Sleep(10 * time.Second)
	for _, js := range jsonic {
		if js.ID == id {
			c.IndentedJSON(http.StatusOK, js)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No message with id : " + id})
}
