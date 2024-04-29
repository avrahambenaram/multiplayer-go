package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/multiplayer-go/internal/config"
	"github.com/avrahambenaram/multiplayer-go/internal/game"
	"github.com/avrahambenaram/multiplayer-go/internal/handlers"
	socketio "github.com/googollee/go-socket.io"
	// "github.com/googollee/go-socket.io/engineio"
	// "github.com/googollee/go-socket.io/engineio/transport"
	// "github.com/googollee/go-socket.io/engineio/transport/polling"
	// "github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
  server := socketio.NewServer(nil)
  myGame := game.New(game.Board{
    Width: 10,
    Height: 10,
  })
  websockets := handlers.NewWebSockets(myGame)

	server.OnConnect("/", websockets.OnConnect)

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		log.Println("chat:", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

  server.OnEvent("/", "movement", websockets.OnMovement)

	server.OnDisconnect("/", websockets.OnDisconnect)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

  http.Handle("/socket.io/", server)
  http.Handle("/", http.FileServer(http.Dir("./public")))

  fmt.Printf("Server running on port %d\n", config.Port)
  http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
