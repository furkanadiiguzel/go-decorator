package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler represents a function that takes an HTTP response writer and request.
type Handler func(http.ResponseWriter, *http.Request)

// Middleware represents a function that takes a Handler and returns a new Handler.
type Middleware func(Handler) Handler

// LoggerMiddleware is a middleware that logs information about incoming requests.
func LoggerMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		defer func() {
			duration := time.Since(startTime)
			fmt.Printf("[%s] %s %s (%v)\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, duration)
		}()

		next(w, r)
	}
}

// AuthMiddleware is a middleware that checks if a request has a valid API key.
func AuthMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != "secret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

// HelloHandler is a simple Handler that responds with a "Hello, world!" message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, world!")
}

// ServeHTTP implements the http.Handler interface for the Handler type.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

/*
func main() {
	// Convert the HelloHandler to a Handler type
	var baseHandler Handler = HelloHandler

	// Apply middlewares to the base handler to create the final handler.
	finalHandler := LoggerMiddleware(AuthMiddleware(baseHandler))

	// Use the final handler as the handler for the root path.
	http.Handle("/", finalHandler)

	// Start the HTTP server on port 8080.
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
*/
