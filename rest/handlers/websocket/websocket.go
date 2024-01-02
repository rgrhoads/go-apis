package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connections = map[string]*websocket.Conn{}

type Request struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMessage(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error: %d", err)
	}

	for {
		var req Request

		err = ws.ReadJSON(&req)
		if err != nil {
			log.Printf("Error: %d", err)
		}

		log.Printf("Message: %v", req.Message)

		connections[req.From] = ws
		req.Message = fmt.Sprintf("%s %s", req.Message, req.To)

		if con, ok := connections[req.To]; ok {
			err := con.WriteJSON(&req)
			if err != nil {
				log.Printf("Error: %d", err)
			}
		}

	}
}
