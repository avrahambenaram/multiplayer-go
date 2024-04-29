package handlers

import (
	"log"
	"math/rand"

	"github.com/avrahambenaram/multiplayer-go/internal/game"
	gameEntity "github.com/avrahambenaram/multiplayer-go/internal/game/entity"
	socketio "github.com/googollee/go-socket.io"
)

type WebSockets struct {
  game *game.Game
}

func NewWebSockets(game *game.Game) *WebSockets {
  return &WebSockets{
    game,
  }
}

func (c *WebSockets) OnConnect(s socketio.Conn) error {
  player := &gameEntity.Player{
    X:  rand.Intn(c.game.Width),
    Y:  rand.Intn(c.game.Height),
    Id: s.ID(),
  }
  c.game.AddPlayer(player)
  log.Println("Connected: ", s.ID())
  return nil
}

func (c *WebSockets) OnMovement(s socketio.Conn, msg string) {
  c.game.MovePlayer(&game.MovePlayerDto{
    PlayerId: s.ID(),
    Direction: msg,
  })
}

func (c *WebSockets) OnDisconnect(s socketio.Conn, reason string) {
  log.Println("Closed: ", reason)
}
