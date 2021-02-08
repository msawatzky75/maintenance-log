package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/msawatzky75/maintenance-log/server/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login contains the options for the login endpoints
type Login struct {
	DB                 *gorm.DB
	JWTSecret          []byte
	AccessTokenCookie  string
	AccessTokenLife    time.Duration
	RefreshTokenCookie string
	RefreshTokenLife   time.Duration
	CookieDomain       string
}

// Error allows for singleton errors
type Error string

func (err Error) Error() string {
	return string(err)
}

// ErrUserNotFound is a singleton error
const ErrUserNotFound = Error("could not find user")

// LoginHandler for login endpoint
func (l *Login) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			l.login(w, r)
		}
	}
}

// AccessTokenClaims is data in the AccessTokenCookie
type AccessTokenClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

// RefreshTokenClaims is data in the RefreshTokenCookie
type RefreshTokenClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}
type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) getUser(c credentials) (*model.User, error) {
	var u model.User
	err := l.DB.Where("email = ?", c.Email).First(&u).Error
	if err != nil {
		return &u, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.Password))
	if err != nil {
		return &u, err
	}

	return &u, nil
}

func (l *Login) login(w http.ResponseWriter, r *http.Request) {
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

	l.SetLoginCookies(u.ID.String(), w)

	json.NewEncoder(w).Encode(map[string]string{"userId": u.ID.String()})
}

// SetLoginCookies sets cookies for a userID
func (l *Login) SetLoginCookies(userID string, w http.ResponseWriter) {
	accessTokenCookie, err := l.createAccessTokenCookie(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, accessTokenCookie)

	refreshTokenCookie, err := l.createRefreshTokenCookie(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, refreshTokenCookie)
}

func (l *Login) createAccessTokenCookie(userID string) (*http.Cookie, error) {
	accessTokenClaims := &AccessTokenClaims{
		UserID: userID,
	}
	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(l.JWTSecret)
	if err != nil {
		return &http.Cookie{}, err
	}

	return &http.Cookie{
		Name:     l.AccessTokenCookie,
		Value:    accessTokenString,
		MaxAge:   int(l.AccessTokenLife.Round(time.Minute).Seconds()),
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") != "DEVELOPMENT",
		Domain:   l.CookieDomain,
		Path:     "/graphql",
		SameSite: http.SameSiteStrictMode,
	}, nil
}

func (l *Login) createRefreshTokenCookie(userID string) (*http.Cookie, error) {
	refreshTokenClaims := &RefreshTokenClaims{
		UserID: userID,
	}
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(l.JWTSecret)
	if err != nil {
		return &http.Cookie{}, err
	}

	return &http.Cookie{
		Name:     l.RefreshTokenCookie,
		Value:    refreshTokenString,
		MaxAge:   int(l.RefreshTokenLife.Round(time.Minute).Seconds()),
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") != "DEVELOPMENT",
		Domain:   l.CookieDomain,
		Path:     "/api/auth/refresh",
		SameSite: http.SameSiteStrictMode,
	}, nil
}
