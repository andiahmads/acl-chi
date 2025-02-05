package main

import (
	"fmt"
	"net/http"
)

// Fungsi ini yang akan di-export dan digunakan oleh Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Menulis respon HTTP
	fmt.Fprintf(w, "Hello from Go on Vercel!")
}

func main() {
	// Menentukan route untuk menangani semua request ke endpoint "/"
	http.HandleFunc("/", Handler)

	// Menggunakan port yang diberikan oleh Vercel
	port := "3000"
	http.ListenAndServe(":"+port, nil)
}

