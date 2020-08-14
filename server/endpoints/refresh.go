package endpoints

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"
)

type refreshParams struct {
	RefreshToken string `json:"refresh_token"`
}

func validateRefreshToken(t string) error {
	return nil
}

// Refresh handles refreshing jwt tokens
func Refresh(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var (
				params refreshParams
				claims jwt.MapClaims
			)
			json.NewDecoder(r.Body).Decode(&params)

			jwt.ParseWithClaims(params.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			uid := claims["userId"].(string)

			err := validateRefreshToken(params.RefreshToken)
			if err != nil {
				http.Error(w, "invalid refresh token", http.StatusBadRequest)
			}
			t, _, err := createTokens(uid)

			json.NewEncoder(w).Encode(map[string]string{
				"access_token":  t,
				"refresh_token": params.RefreshToken,
			})
		}
	}
}
