package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", Handler)

	log.Println("Server running on http://localhost:9999")
	http.ListenAndServe(":9999", r)
}
