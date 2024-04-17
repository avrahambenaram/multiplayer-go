package game

import (
	"math/rand"
	"time"
)

type Board struct {
  Width  int `json:"width"`
  Height int `json:"height"`
}

type Game struct {
  *Board
  fruits    []*Fruit
  players   []*Player
  movements map[string]func(player *Player)
  stop      chan bool
}

type MovePlayerDto struct {
  PlayerId  string `json:"playerId"`
  Direction string `json:"direction"`
}

type StartGameDto struct {
  GenerateFruitsInSeconds int
  MaxPoints               int
}

func New() *Game {
  game := &Game{
    fruits: make([]*Fruit, 0, 10),
    players: make([]*Player, 0, 10),
    movements: make(map[string]func(player *Player), 4),
  }
  game.movements["left"] = game.movePlayerLeft
  game.movements["right"] = game.movePlayerRight
  game.movements["up"] = game.movePlayerUp
  game.movements["down"] = game.movePlayerDown
  return game
}

func (c *Game) AddPlayer(player *Player) {
  c.players = append(c.players, player)
}

func (c *Game) RemovePlayer(playerId string) {
  for i, player := range(c.players) {
    if player.Id == playerId {
      c.players = append(c.players[:i], c.players[i+1:]...)
    }
  }
}

func (c *Game) MovePlayer(props *MovePlayerDto) {
  player := c.FindPlayerById(props.PlayerId)
  if player == nil {
    return
  }

  movePlayer := c.movements[props.Direction]
  if movePlayer != nil {
    movePlayer(player)
  }
}

func (c *Game) movePlayerLeft(player *Player) {
  if player.X > 0 {
    player.X--
  }
}

func (c *Game) movePlayerRight(player *Player) {
  if player.X < c.Width {
    player.X++
  }
}

func (c *Game) movePlayerUp(player *Player) {
  if player.Y > 0 {
    player.Y--
  }
}

func (c *Game) movePlayerDown(player *Player) {
  if player.Y < c.Height {
    player.Y++
  }
}

func (c *Game) FindPlayerById(playerId string) *Player {
  for _, player := range(c.players) {
    if player.Id == playerId {
      return player
    }
  }
  return nil
}

func (c *Game) Start(props StartGameDto) {
  ticker := time.NewTicker(time.Duration(props.GenerateFruitsInSeconds))
  c.stop = make(chan bool)

  go func() {
    for {
      select {
      case <- ticker.C:
        c.generateFruit()
      case <- c.stop:
        ticker.Stop()
        close(c.stop)
        c.stop = nil
        return
      }
    }
  }()

  <- c.stop
}

func (c *Game) generateFruit() {
  var fruitType string
  x, y := rand.Intn(c.Width), rand.Intn(c.Height)

  if rand.Intn(3) > 0 {
    fruitType = "special"
  } else {
    fruitType = "normal"
  }

  fruit := &Fruit{
    x,
    y,
    fruitType,
  }
  c.fruits = append(c.fruits, fruit)
}

func (c *Game) Stop() {
  c.stop <- true
}
