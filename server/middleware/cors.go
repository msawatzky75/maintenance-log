package middleware

import (
	"net/http"
)

// Cors is the configuration for the cors middleware
type Cors struct {
	Cors string
}

// Handler is the cors middleware handler
func (c *Cors) Handler(h http.Handler, methods string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", c.Cors)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie")
		w.Header().Set("Vary", "Origin")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", methods+", OPTIONS")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
