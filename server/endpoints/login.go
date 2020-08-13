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

func validateUser(db *gorm.DB, email string, password string) (u model.User, e error) {
	e = db.Where("email = ?", email).First(&u).Error
	if e != nil {
		return
	}

	e = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
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
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return
	}

	rt, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 7).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	return
}

// Login enpoint handler
func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := validateUser(db, email, password)
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
}
