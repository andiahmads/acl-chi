package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var rolesPermissions = map[string][]string{
	"user":  {"read", "read-product"},
	"admin": {"read", "write"},
}

func ACLMiddleware(permission []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
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

func main() {
	r := chi.NewRouter()

	r.With(ACLMiddleware([]string{"read"})).Get("/read", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read operation allowed"))
	})

	r.With(ACLMiddleware([]string{"read-product"})).Get("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Read product done"))
	})

	r.With(ACLMiddleware([]string{"write"})).Get("/write", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Write operation allowed"))
	})

	http.ListenAndServe(":9999", r)

}
