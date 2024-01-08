package games

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Action struct {
	Action   string `json:"action"`
	Data     string `json:"data"`
	GameId   string `json:"gameId"`
	Username string `json:"username"`
}

type Board struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}

type Game struct {
	// Board   Board           `json:"board"`
	Id      string          `json:"id"`
	Board   [10]string      `json:"board"`
	ConnX   *websocket.Conn `json:"connX"`
	ConnO   *websocket.Conn `json:"connO"`
	PlayerX string          `json:"pX"`
	PlayerO string          `json:"pO"`
}

type Move struct {
	Username string `json:"username"`
	Move     string `json:"move"`
}

// var connections = map[string]*websocket.Conn{}

var games = map[string]*Game{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func TicTacToe(c *gin.Context) {
	gameType := c.Param("gameType")
	log.Printf("Game Type: %s", gameType)

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error: %d", err)
	}

	if len(games) == 0 {
		var id = uuid.NewString()
		var game Game

		games[id] = &game
	}

	for {
		var action Action

		err = ws.ReadJSON(&action)
		if err != nil {
			log.Printf("Error: %d", err)
		}

		log.Println(action.Username)
		// connections[action.Username] = ws

		if action.Action == "connect" {
			log.Println("Looking for Opponent")

			for id, game := range games {
				var gameId string

				if game.PlayerX == "" {
					log.Println("Player X is empty")

					game.Id = id

					game.PlayerX = action.Username
					game.ConnX = ws

					err = game.ConnX.WriteJSON(game)
					if err != nil {
						log.Printf("Error: %d", err)
					}
				} else if game.PlayerO == "" {
					log.Println("Player O is empty")

					game.Id = id

					game.PlayerO = action.Username
					game.ConnO = ws

					err = game.ConnX.WriteJSON(game)
					if err != nil {
						log.Printf("Error: %d", err)
					}

					err = game.ConnO.WriteJSON(game)
					if err != nil {
						log.Printf("Error: %d", err)
					}
				} else {
					log.Println("Creating new game")
					var game Game

					gameId = uuid.NewString()
					game.Id = gameId

					games[gameId] = &game
					game.PlayerX = action.Username
					game.ConnX = ws

					err = game.ConnX.WriteJSON(game)
					if err != nil {
						log.Printf("Error: %d", err)
					}
				}
			}
		} else if action.Action == "move" {
			log.Printf("Player %s selected board position %s", action.Username, action.Data)
			log.Printf("Game ID: %s", action.GameId)

			updatedGame := games[action.GameId]
			position, err := strconv.Atoi(action.Data)
			if err != nil {
				log.Printf("Error: %d", err)
			}

			if updatedGame.PlayerX == action.Username {
				updatedGame.Board[position] = "X"
			} else if updatedGame.PlayerO == action.Username {
				updatedGame.Board[position] = "O"
			}

			err = games[action.GameId].ConnX.WriteJSON(&updatedGame)
			if err != nil {
				log.Printf("Error: %d", err)
			}

			err = games[action.GameId].ConnO.WriteJSON(&updatedGame)
			if err != nil {
				log.Printf("Error: %d", err)
			}

		}

		// else {
		// 	if con, ok := connections[action.Username]; ok {
		// 		err := con.WriteJSON(&action)
		// 		if err != nil {
		// 			log.Printf("Error: %d", err)
		// 		}
		// 	}
		// }

		// log.Println(action.Username)
		// connections[action.Username] = ws

	}
}

// func SubmitMove(c *gin.Context) {
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Printf("Error: %d", err)
// 	}

// 	for {
// 		var move Move

// 		err = ws.ReadJSON(&move)
// 		if err != nil {
// 			log.Printf("Error: %d", err)
// 		}

// 		log.Println(move)

// 		connections[move.Username] = ws

// 		if con, ok := connections[move.Username]; ok {
// 			err := con.WriteJSON(&move)
// 			if err != nil {
// 				log.Printf("Error: %d", err)
// 			}
// 		}

// 	}
// }
