package game

import (
	"math/rand"
	"time"

	"github.com/avrahambenaram/multiplayer-go/game/entity"
)

type StartGameDto struct {
  GenerateFruitsInSeconds int
  MaxPoints               int
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
  defer c.fruitMutex.Unlock()

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

  fruit := &entity.Fruit{
    X: x,
    Y: y,
    Type: fruitType,
  }
  c.fruits[x] = append(c.fruits[x], fruit)
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
