package main

import (
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Hello World!\n")
}

func main() {
  http.HandleFunc("/", hello)

  fmt.Printf("Listening on port %d\n", 3000)
  http.ListenAndServe(":3000", nil)
}
