package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/msawatzky75/maintenance-log/server/graph/model"
)

// RefreshHandler handles refreshing jwt tokens
func (l *Login) RefreshHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			l.refresh(w, r)
		}
	}
}

func (l *Login) refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(l.RefreshTokenCookie)
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("refresh: Cookie not found")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("refresh: Error reading cookie")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Initialize a new instance of `Claims`
	claims := &RefreshTokenClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	t, err := jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return l.JWTSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("refresh: Invalid Signature")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("refresh: Error parsing refresh token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !t.Valid {
		log.Println("refresh: Invalid Token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var u model.User
	err = l.DB.Where("id = ?", claims.UserID).Find(&u).Error
	if err != nil {
		log.Println("refresh: Unable to find user")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessTokenCookie, err := l.createAccessTokenCookie(u.ID.String())
	if err != nil {
		log.Println("refresh: Unable to create cookie")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, accessTokenCookie)

	json.NewEncoder(w).Encode(map[string]string{"userId": u.ID.String()})
}
