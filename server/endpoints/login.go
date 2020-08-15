package endpoints

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	DB                 *gorm.DB
	JWTSecret          []byte
	AccessTokenCookie  string
	AccessTokenLife    time.Duration
	RefreshTokenCookie string
	RefreshTokenLife   time.Duration
}

// Handler for login endpoint
func (l *Login) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			l.signin(w, r)
		}
	}
}

type AccessTokenClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}
type RefreshTokenClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}
type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) getUser(c credentials) (u *model.User, e error) {
	e = l.DB.Where("email = ?", c.Email).First(u).Error
	if e != nil {
		return
	}

	e = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.Password))
	if e != nil {
		return
	}

	return
}

// Create the Signin handler
func (l *Login) signin(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := l.getUser(creds)

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

	refreshTokenExiprationTime := time.Now().Add(5 * time.Minute)
	refreshTokenClaims := &RefreshTokenClaims{
		UserID: u.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExiprationTime.Unix(),
		},
	}
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(l.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer http.SetCookie(w, &http.Cookie{
		Name:     l.RefreshTokenCookie,
		Value:    refreshTokenString,
		Expires:  refreshTokenExiprationTime,
		HttpOnly: true,
		Secure:   true,
	})
}
