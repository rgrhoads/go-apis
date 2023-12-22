package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Request struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
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

		log.Println(req)
	}

	// ws.WriteJSON()
}
