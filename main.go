package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

// ACLMiddleware to check permissions
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

// main handler function
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Vercel!")
}

// Vercel entry point function
func Handler(w http.ResponseWriter, r *http.Request) {
	routes := chi.NewRouter()
	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)

	routes.Get("/", mainHandler)

	// Apply ACL middleware for routes with permission checks
	routes.With(ACLMiddleware([]string{"read"})).Get("/read", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read operation allowed"))
	})

	routes.With(ACLMiddleware([]string{"read-product"})).Get("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read product done"))
	})

	routes.With(ACLMiddleware([]string{"write"})).Get("/write", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Write operation allowed"))
	})

	http.Handle("/", routes)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

