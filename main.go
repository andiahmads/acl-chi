package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var rolesPermissions = map[string][]string{
	"user":  {"read", "read-product"},
	"admin": {"read", "write"},
}

func ACLMiddleware(permission []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := "user"

			requiredPermissions := make(map[string]bool)
			for _, p := range permission {
				requiredPermissions[p] = true
			}

			userPermission, exist := rolesPermissions[userRole]
			if !exist {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			var authorized bool
			for _, p := range userPermission {
				fmt.Println("list permission", p)
				if requiredPermissions[p] {
					authorized = true
					break
				}
			}

			if !authorized {
				http.Error(w, "Forbiden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})

	}
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Read operation allowed"))

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	http.HandleFunc("/read", ReadHandler)

	r.With(ACLMiddleware([]string{"read-product"})).Get("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read product done"))
	})

	r.With(ACLMiddleware([]string{"write"})).Get("/write", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Write operation allowed"))
	})

	port := 9999
	log.Printf("server running on port:%d", port)
	http.ListenAndServe(":9999", r)

}
