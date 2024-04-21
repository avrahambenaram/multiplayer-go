package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/avrahambenaram/multiplayer-go/internal/config"
)

func index(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  header.Set("Content-Type", "text/html")

  file, fileErr := os.ReadFile("./public/index.html")
  if fileErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("An error occurred"))
    return
  }
  w.Write(file)
}
func main() {
  http.HandleFunc("/", index)

  fmt.Printf("Server running on port %d\n", config.Port)
  http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
