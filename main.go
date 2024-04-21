package main

import (
	"fmt"
	"net/http"

	"github.com/avrahambenaram/multiplayer-go/internal/config"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("./public")))

  fmt.Printf("Server running on port %d\n", config.Port)
  http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
