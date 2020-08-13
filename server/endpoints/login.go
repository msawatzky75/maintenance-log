package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
	"golang.org/x/crypto/bcrypt"
)

type authClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}
type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func validateUser(db *gorm.DB, l login) (u model.User, e error) {
	e = db.Where("email = ?", l.Email).First(&u).Error
	if e != nil {
		return
	}

	e = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
	if e != nil {
		return
	}

	return
}

func createTokens(userID string) (t string, rt string, err error) {
	// Generate encoded token and send it as response.
	t, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &authClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return
	}

	rt, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &authClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 7).Unix(),
		},
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	return
}

func loginPOST(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var l login

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		panic(err)
	}

	user, err := validateUser(db, l)
	if err != nil {
		http.Error(w, "unauthorized", 401)
		return
	}

	t, rt, err := createTokens(user.ID.String())

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	})
}

// Login enpoint handler
func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			loginPOST(db, w, r)
		}
	}
}
