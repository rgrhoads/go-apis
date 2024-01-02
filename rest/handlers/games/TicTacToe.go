package games

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Board struct {
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
	Four  string `json:"four"`
	Five  string `json:"five"`
	Six   string `json:"six"`
	Seven string `json:"seven"`
	Eight string `json:"eight"`
	Nine  string `json:"nine"`
}

type User struct {
	Username string `json:"username"`
}

type Game struct {
	Id      string `json:"id"`
	Board   Board  `json:"board"`
	PlayerX string `json:"pX"`
	PlayerO string `json:"pO"`
}

type Move struct {
	Username string `json:"username"`
	Move     string `json:"move"`
}

var connections = map[string]*websocket.Conn{}

var games = map[string]*Game{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Connect(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error: %d", err)
	}

	for {
		var user User

		err = ws.ReadJSON(&user)
		if err != nil {
			log.Printf("Error: %d", err)
		}

		log.Println(user.Username)

		connections[user.Username] = ws

		if con, ok := connections[user.Username]; ok {
			err := con.WriteJSON(&user)
			if err != nil {
				log.Printf("Error: %d", err)
			}
		}
	}
}

func SubmitMove(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error: %d", err)
	}

	for {
		var move Move

		err = ws.ReadJSON(&move)
		if err != nil {
			log.Printf("Error: %d", err)
		}

		log.Println(move)

		connections[move.Username] = ws

		if con, ok := connections[move.Username]; ok {
			err := con.WriteJSON(&move)
			if err != nil {
				log.Printf("Error: %d", err)
			}
		}

	}
}
