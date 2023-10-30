package main

import (
	"example/main/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Log in and Sign up
	router.POST("/registration", app.Registration)
	router.GET("/authorization", app.Authorization)

	//Websockets
	router.GET("/get_message", app.GetMessageByID)
	//Account statistics
	router.GET("/is_friends", app.IsFriends)
	router.Run("localhost:8000")
}
