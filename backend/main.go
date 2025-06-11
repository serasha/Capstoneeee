package main

import (
  "net/http"
  "fmt"
)

func main() {
  fs := http.FileServer(http.Dir("../frontend/dist"))
  http.Handle("/", fs)
  http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"pimpin kontol"}`))
  })

  fmt.Println("Listening at http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}
