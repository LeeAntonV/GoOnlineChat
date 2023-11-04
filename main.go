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
	router.GET("/get_profile", app.GetProfile)

	//Websockets
	router.GET("/get_message", app.GetMessageByID)
	//Account statistics
	router.Run("localhost:8000")
}
