package main // This is absolutely required

import (
	"fmt"
	"net/http" // If you're building a web server
)

func main() { // This is the entry point
	fmt.Println("Hello from Vercel!")

	// Example for a web server:
	http.HandleFunc("/", handler) // Define your handler function
	port := "9999"                // or get it from an environment variable
	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

