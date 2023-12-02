package main

import (
	"example/main/internal/handlers"
	"github.com/gin-gonic/gin"
)

// TODO: cookies, websocket, file loading into db
func main() {
	router := gin.Default()
	//Log in and Sign up
	router.POST("/registration", handlers.Registration)
	router.GET("/authorization", handlers.Authorization)
	router.GET("/get_profile", handlers.GetProfile)
	router.GET("/auth", handlers.AuthMiddleware)

	//Websockets
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/get_message", handlers.GetMessageByID)
	router.GET("/ws", handlers.WsHandler)

	//Account statistics
	router.Run("localhost:8000")
}
