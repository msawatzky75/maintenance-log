package endpoints

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

// Refresh handles refreshing jwt tokens
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
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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

	var u model.User
	err = l.DB.Where("id = ?", claims.UserID).Find(&u).Error
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessTokenExiprationTime := time.Now().Add(5 * time.Minute)
	accessTokenClaims := &AccessTokenClaims{
		UserID: u.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExiprationTime.Unix(),
		},
	}
	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(l.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer http.SetCookie(w, &http.Cookie{
		Name:     l.AccessTokenCookie,
		Value:    accessTokenString,
		Expires:  accessTokenExiprationTime,
		HttpOnly: true,
		Secure:   true,
	})
}
