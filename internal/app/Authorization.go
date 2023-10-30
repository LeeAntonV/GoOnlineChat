package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pattern struct {
	ID        string `json:"id"`
	Name      string `json:"Name"`
	Password  string `json:"password"`
	IsFriends bool   `json:"IsFriends"`
}

var ProfileAuth = []Pattern{
	{ID: "1", Name: "Anton Lee", Password: "1234", IsFriends: true},
	{ID: "2", Name: "Maksym Stashuk", Password: "1234", IsFriends: true},
}

func Authorization(c *gin.Context) {
	Name := c.GetHeader("name")
	Password := c.GetHeader("password")
	fmt.Println(Name + Password)
	for _, prof := range ProfileAuth {
		if prof.Name == Name && prof.Password == Password {
			c.IndentedJSON(http.StatusAccepted, prof)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "WrongCredentials"})
}
