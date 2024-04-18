package game

import (
	"math/rand"
	"sync"
	"time"
)

type Board struct {
  Width  int `json:"width"`
  Height int `json:"height"`
}

type Game struct {
  *Board
  players    []*Player
  fruits     map[int][]*Fruit
  movements  map[string]func(player *Player)
  maxPoints  int
  stop       chan bool
  fruitMutex sync.Mutex
}

type MovePlayerDto struct {
  PlayerId  string `json:"playerId"`
  Direction string `json:"direction"`
}

type StartGameDto struct {
  GenerateFruitsInSeconds int
  MaxPoints               int
}

func New(board Board) *Game {
  game := &Game{
    Board: &board,
    fruits: make(map[int][]*Fruit, board.Width),
    players: make([]*Player, 0, 10),
    movements: make(map[string]func(player *Player), 4),
  }
  game.movements["left"] = game.movePlayerLeft
  game.movements["right"] = game.movePlayerRight
  game.movements["up"] = game.movePlayerUp
  game.movements["down"] = game.movePlayerDown

  for i := 0; i < board.Width; i++ {
    game.fruits[i] = make([]*Fruit, 0, board.Height)
  }
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
    c.checkPlayerCollision(player)
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

func (c *Game) checkPlayerCollision(player *Player) {
  fruits := c.fruits[player.X]
  for i, fruit := range(fruits) {
    if fruit.Y == player.Y {
      fruits = append(fruits[i:], fruits[:i]...)
      if fruit.Type == "special" {
        player.Points += 5
      }
      if fruit.Type == "normal" {
        player.Points += 1
      }
      if player.Points == c.maxPoints {
        c.Stop()
      }
    }
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
  c.maxPoints = props.MaxPoints
  ticker := time.NewTicker(time.Duration(props.GenerateFruitsInSeconds))
  c.stop = make(chan bool)

  go func() {
    for {
      select {
      case <- ticker.C:
        go c.generateFruit()
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
  c.fruitMutex.Lock()
  var fruitType string
  canBePlaced := false
  x, y := rand.Intn(c.Width), rand.Intn(c.Height)

  for len(c.fruits[x]) == 10 {
    x = rand.Intn(c.Width)
  }

  outerLoop:
  for !canBePlaced {
    for _, fruit := range(c.fruits[x]) {
      if fruit.Y == y {
        y = rand.Intn(c.Height)
        break outerLoop
      }
    }
  }

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
  c.fruits[x] = append(c.fruits[x], fruit)
  c.fruitMutex.Unlock()
}

func (c *Game) Stop() {
  c.stop <- true
  c.resetPoints()
}

func (c *Game) resetPoints() {
  for _, player := range(c.players) {
    player.Points = 0
  }
}
