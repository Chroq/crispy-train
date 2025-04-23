package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the API!")
	})

	healthHandler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "OK")
		},
	)
	http.HandleFunc("/api/v1/health", logMiddleware(healthHandler).ServeHTTP)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeFormated := time.Now().Format(time.RFC1123)
		fmt.Printf("Received request: %s %s at %s\n", r.Method, r.URL.Path, timeFormated)
		next.ServeHTTP(w, r)
	})
}
