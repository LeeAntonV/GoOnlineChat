package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfilePattern struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
}

func Registration(c *gin.Context) {
	var newJson ProfilePattern

	if err := c.BindJSON(&newJson); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, newJson)
	fmt.Println(newJson)
}
