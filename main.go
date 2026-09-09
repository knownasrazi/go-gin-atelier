package main

import "net/http"

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("go-gin-atelier - Gin, atelier-grade."))
  })
  http.ListenAndServe(":8080", nil)
}
