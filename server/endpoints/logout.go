package endpoints

import (
	"net/http"
)

// LogoutHandler deletes both AccessTokenCookie and RefreshTokenCookie
func (l *Login) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		http.SetCookie(w, &http.Cookie{Name: l.AccessTokenCookie, MaxAge: -1})
		http.SetCookie(w, &http.Cookie{Name: l.RefreshTokenCookie, MaxAge: -1})
	}
}
