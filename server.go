package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
//	"encoding/json"

	"github.com/satori/go.uuid"
	"github.com/googollee/go-socket.io"
)

// Event messages
const (
	EVENT_CHAT_MESSAGE = "chat message" 
	EVENT_CONNECT = "connection" 
	EVENT_DISCONNECT = "disconnection"
	EVENT_ERROR  = "error"
	EVENT_START_GAME = "start game"
)

// Game related struct
type Game struct {
	Game_id uuid.UUID `json:"game_id"`
	Created_at time.Time `json:"created_at"`
}

func main() {
	server, err := socketio.NewServer(nil)
	var room_id string;

	if err != nil {
		log.Fatal(err)
	}

	server.On(EVENT_CONNECT, func(so socketio.Socket) {
		log.Println("on connection venkat")

		so.On(EVENT_START_GAME, func(msg string) Game {
			log.Println(EVENT_START_GAME + "msg: "+ msg)

			// error handling
			u2, err := uuid.NewV4()
			if err != nil {
				fmt.Printf("Something went wrong: %s", err)
				panic(err)
			}

			g := Game{ Game_id: u2, Created_at: time.Now().Local()}
			room_id = u2.String()
			if err != nil {
				fmt.Printf("Something went wrong: %s", err)
				panic (err)
			}

			// join the client in the game room
			so.Join(room_id)
			return g
		})

		so.On(EVENT_CHAT_MESSAGE, func(msg string) {
			log.Println(EVENT_CHAT_MESSAGE)
			so.Emit(EVENT_CHAT_MESSAGE, "dei venna")
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