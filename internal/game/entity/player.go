package entity

type Player struct {
  X      int    `json:"x"`
  Y      int    `json:"y"`
  Id     string `json:"id"`
  Points int    `json:"points"`
}
