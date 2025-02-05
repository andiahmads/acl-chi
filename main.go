package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// ACLMiddleware untuk memeriksa izin akses
func ACLMiddleware(permission []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := "user"
			requiredPermissions := make(map[string]bool)

			for _, p := range permission {
				requiredPermissions[p] = true
			}

			rolesPermissions := map[string][]string{
				"user":  {"read", "read-product"},
				"admin": {"read", "write"},
			}

			userPermission, exist := rolesPermissions[userRole]
			if !exist {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			var authorized bool
			for _, p := range userPermission {
				if requiredPermissions[p] {
					authorized = true
					break
				}
			}

			if !authorized {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Fungsi ini akan menangani rute / untuk aplikasi
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Vercel!")
}

// Fungsi entry point yang harus diekspor untuk Vercel
func main() {
	routes := chi.NewRouter()
	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)

	// Rute utama
	routes.Get("/", Handler)

	// Rute dengan middleware ACL
	routes.With(ACLMiddleware([]string{"read"})).Get("/read", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read operation allowed"))
	})

	routes.With(ACLMiddleware([]string{"read-product"})).Get("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read product done"))
	})

	routes.With(ACLMiddleware([]string{"write"})).Get("/write", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Write operation allowed"))
	})

	port := "9999"
	// Menjalankan server
	http.Handle("/", routes)
	http.ListenAndServe(":"+port, nil)
}

