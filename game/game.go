package game

type Board struct {
  Width  int `json:"width"`
  Height int `json:"height"`
}

type Game struct {
  *Board
  fruits   []*Fruit
  players  []*Player
  movements map[string]func(player *Player)
}

type MovePlayerDto struct {
  PlayerId  string `json:"playerId"`
  Direction string `json:"direction"`
}

type StartGameDto struct {
  GenerateFruitsInSeconds int
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
