package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients []websocket.Conn

func WsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()
	clients = append(clients, *conn)
	if err != nil {
		log.Println(err.Error())
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println(err.Error())
				break
			}
			break
		}

		for _, client := range clients {
			if err = client.WriteMessage(messageType, message); err != nil {
				log.Println(err.Error())
			}
		}
	}
}
