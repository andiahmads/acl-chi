package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rctx := chi.NewRouter()
	rctx.Use(middleware.RequestID)
	rctx.Use(middleware.RealIP)
	rctx.Use(middleware.Logger)
	rctx.Use(middleware.Recoverer)

	rctx.Get("/read", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read operation allowed"))
	})

	rctx.Get("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read product done"))
	})

	rctx.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server started on port 3000")
	http.ListenAndServe(":3000", nil)
}

