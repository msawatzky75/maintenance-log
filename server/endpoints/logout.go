package endpoints

import (
	"net/http"
)

// Handler for login endpoint
func (l *Login) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		http.SetCookie(w, &http.Cookie{Name: l.AccessTokenCookie, MaxAge: -1})
	}
}
