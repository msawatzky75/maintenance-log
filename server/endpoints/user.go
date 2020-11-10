package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/msawatzky75/maintenance-log/server/graph/model"
)

func (l *Login) signup(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = l.getUser(creds)
	// This is the error we are looking for
	if err != ErrUserNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u, err := model.CreateUser(l.DB, creds.Email, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// User has been created, log them in
	accessTokenCookie, err := l.createAccessTokenCookie(u.ID.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer http.SetCookie(w, accessTokenCookie)

	refreshTokenCookie, err := l.createRefreshTokenCookie(u.ID.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer http.SetCookie(w, refreshTokenCookie)
}

// User is the user endpoint handler
func (l *Login) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			l.signup(w, r)
		default:
			fmt.Fprint(w, "Only POST is supported for this endpoint")
		}
	}
}
