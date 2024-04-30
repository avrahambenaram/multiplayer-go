package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/multiplayer-go/internal/config"
	// "github.com/avrahambenaram/multiplayer-go/internal/game"
	// "github.com/avrahambenaram/multiplayer-go/internal/handlers"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
  server := socketio.NewServer(nil)
  // myGame := game.New(game.Board{
  //   Width: 10,
  //   Height: 10,
  // })
  // websockets := handlers.NewWebSockets(myGame)

	server.OnConnect("/", func(s socketio.Conn) error {
    s.SetContext("")
    log.Println("Connected: ", s.ID())

    return nil
  })
  server.OnEvent("/", "ping", func(s socketio.Conn, msg string) {
    s.Emit("pong", "hello")
  })

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
    log.Println("Closed: ", reason)
  })

  go server.Serve()
	defer server.Close()

  http.Handle("/socket.io/", server)
  http.Handle("/", http.FileServer(http.Dir("./public")))

  fmt.Printf("Server running on port %d\n", config.Port)
  http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
