package main

import (
	"log"
	"net/http"
//	"time"
//	"fmt"

	"./gameDetails"
	"./store"
	//"./util"
//	"github.com/satori/go.uuid"
	"github.com/googollee/go-socket.io"
)

// Event messages
const (
	EVENT_CHAT_MESSAGE = "chat message" 
	EVENT_CONNECT = "connection" 
	EVENT_DISCONNECT = "disconnection"
	EVENT_ERROR  = "error"
	EVENT_JOIN_GAME = "join game"
	EVENT_START_GAME = "start game"
	EVENT_PLAYER_STRIKE= "player strike"
)

func main() {
	server, err := socketio.NewServer(nil)
	store.GameStore.LocalStore *gameStore := &store.LocalGameStore{}
	var room_id string;

	if err != nil {
		log.Fatal(err)
	}

	server.On(EVENT_CONNECT, func(so socketio.Socket) {
		log.Println(EVENT_CONNECT + " sockId: " + so.Id())
		player := gameDetails.NewPlayer(so.Id())
		
		so.On(EVENT_JOIN_GAME, func(msg gameDetails.GameBoard) gameDetails.Game {
			log.Println(EVENT_JOIN_GAME + " Msg: ")
			
			player.FillPlayerBoard(&msg)
			log.Println(player.Board)
			game := gameStore.Allocate(player)

			// join the client in the game room
			so.Join(game.Game_id)
			log.Println(game)
			return *game
		})

		so.On(EVENT_CHAT_MESSAGE, func(msg string) {
			log.Println(EVENT_CHAT_MESSAGE)
			so.Emit(EVENT_CHAT_MESSAGE, "xxxx")
			so.BroadcastTo(room_id, EVENT_CHAT_MESSAGE, msg)
		})

		so.On(EVENT_DISCONNECT, func() {
			log.Println("on disconnect")
		})
	})

	server.On(EVENT_ERROR, func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}