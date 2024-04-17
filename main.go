package main

import (
	"fmt"
	"net/http"
	"os"
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
  fmt.Println("Server running on port 3000")
  http.ListenAndServe(":3000", nil)
}
