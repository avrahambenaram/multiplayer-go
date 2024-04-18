package game

import (
	"sync"

	"github.com/avrahambenaram/multiplayer-go/game/entity"
)

type Board struct {
  Width  int `json:"width"`
  Height int `json:"height"`
}

type Game struct {
  *Board
  Players    []*entity.Player
  Fruits     map[int][]*entity.Fruit
  movements  map[string]func(player *entity.Player)
  maxPoints  int
  running    bool
  stop       chan bool
  fruitMutex sync.Mutex
}

func New(board Board) *Game {
  game := &Game{
    Board: &board,
    Fruits: make(map[int][]*entity.Fruit, board.Width),
    Players: make([]*entity.Player, 0, 10),
    movements: make(map[string]func(player *entity.Player), 4),
  }
  game.movements["left"] = game.movePlayerLeft
  game.movements["right"] = game.movePlayerRight
  game.movements["up"] = game.movePlayerUp
  game.movements["down"] = game.movePlayerDown

  for i := 0; i < board.Width; i++ {
    game.Fruits[i] = make([]*entity.Fruit, 0, board.Height)
  }
  return game
}
