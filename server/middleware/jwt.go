package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const JwtContextKey string = "user"

type Jwt struct {
	AccessTokenCookie  string
	RefreshTokenCookie string
	Secret             []byte
}

// Handler is the Jwt middleware handler
func (j *Jwt) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie(j.AccessTokenCookie)
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)

				if r.Header.Get("content-type") == "application/json" {
					json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized/No Cookie"})
				}
				log.Println("No cookie :(")
				return
			} else {
				// For any other type of error, return a bad request status
				w.WriteHeader(http.StatusBadRequest)
				log.Println("Failed to read cookie", err)
			}
			return
		}

		// Initialize a new instance of `Claims`
		claims := &jwt.MapClaims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		t, err := jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return j.Secret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !t.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), JwtContextKey, claims)

		log.Println("Cookie found and accepted, continuing...")
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
