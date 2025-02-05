package main

import (
	"chi-acl/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", handler.Handler)

	log.Println("Server running on http://localhost:9999")
	http.ListenAndServe(":9999", r)
}
