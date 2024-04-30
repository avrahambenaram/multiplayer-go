package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/multiplayer-go/internal/config"
	"github.com/avrahambenaram/multiplayer-go/internal/game"
	"github.com/avrahambenaram/multiplayer-go/internal/handlers"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
  server := socketio.NewServer(nil)
  myGame := game.New(game.Board{
    Width: 10,
    Height: 10,
  })
  websockets := handlers.NewWebSockets(myGame)

	server.OnConnect("/", websockets.OnConnect)
  server.OnEvent("/", "movement", websockets.OnMovement)

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", websockets.OnDisconnect)

  go server.Serve()
	defer server.Close()

  http.Handle("/socket.io/", server)
  http.Handle("/", http.FileServer(http.Dir("./public")))

  fmt.Printf("Server running on port %d\n", config.Port)
  http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
