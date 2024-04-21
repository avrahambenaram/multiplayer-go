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
  ticker := time.NewTicker(time.Duration(props.GenerateFruitsInSeconds)*time.Second)
  c.running = true
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
}

func (c *Game) generateFruit() {
  c.fruitMutex.Lock()
  defer c.fruitMutex.Unlock()

  var fruitType string
  canBePlaced := false
  x, y := rand.Intn(c.Width), rand.Intn(c.Height)

  for len(c.Fruits[x]) == 10 {
    x = rand.Intn(c.Width)
  }

  outerLoop:
  for !canBePlaced {
    if len(c.Fruits[x]) == 0 {
      canBePlaced = true
    }

    for i, fruit := range(c.Fruits[x]) {
      if fruit.Y == y {
        y = rand.Intn(c.Height)
        break outerLoop
      } else if i == len(c.Fruits[x]) - 1 {
        canBePlaced = true
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

  if c.running {
    c.Fruits[x] = append(c.Fruits[x], fruit)
  }
}

func (c *Game) Stop() {
  c.running = false
  c.stop <- true
  c.resetPoints()
  c.cleanFruits()
}

func (c *Game) resetPoints() {
  for _, player := range(c.Players) {
    player.Points = 0
  }
}
