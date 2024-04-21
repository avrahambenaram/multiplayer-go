package config

import "github.com/spf13/viper"

type gameConfig struct {
  Width  int
  Height int
}

var (
  Port int
  Game *gameConfig
)

func init() {
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")
  viper.SetDefault("port", 3000)
  viper.SetDefault("game.width", 10)
  viper.SetDefault("game.height", 10)
  err := viper.ReadInConfig()
  if err != nil {
    panic(err)
  }

  Port = viper.GetInt("port")
  Game = &gameConfig{}
  Game.Width = viper.GetInt("game.width")
  Game.Height = viper.GetInt("game.height")
}
