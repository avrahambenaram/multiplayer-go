package game

import (
	"slices"

	"github.com/avrahambenaram/multiplayer-go/internal/game/entity"
)

type MovePlayerDto struct {
  PlayerId  string `json:"playerId"`
  Direction string `json:"direction"`
}

func (c *Game) AddPlayer(player *entity.Player) {
  c.Players = append(c.Players, player)
}

func (c *Game) RemovePlayer(playerId string) {
  for i, player := range(c.Players) {
    if player.Id == playerId {
      c.Players = slices.Delete(c.Players, i, i+1)
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

func (c *Game) movePlayerLeft(player *entity.Player) {
  if player.X > 0 {
    player.X--
  }
}

func (c *Game) movePlayerRight(player *entity.Player) {
  if player.X < c.Width {
    player.X++
  }
}

func (c *Game) movePlayerUp(player *entity.Player) {
  if player.Y > 0 {
    player.Y--
  }
}

func (c *Game) movePlayerDown(player *entity.Player) {
  if player.Y < c.Height {
    player.Y++
  }
}

func (c *Game) checkPlayerCollision(player *entity.Player) {
  fruits := c.Fruits[player.X]
  for i, fruit := range(fruits) {
    if fruit.Y == player.Y {
      fruits = slices.Delete(fruits, i, i+1)
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

func (c *Game) FindPlayerById(playerId string) *entity.Player {
  for _, player := range(c.Players) {
    if player.Id == playerId {
      return player
    }
  }
  return nil
}
