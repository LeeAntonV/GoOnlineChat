package main

import (
	app2 "example/main/cmd/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Log in and Sign up
	router.POST("/registration", app2.Registration)
	router.GET("/authorization", app2.Authorization)

	//Websockets
	router.GET("/get_message", app2.GetMessageByID)
	//Account statistics
	router.GET("/is_friends", app2.IsFriends)
	router.Run("localhost:8000")
}
