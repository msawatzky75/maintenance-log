package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

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
			var params refreshParams
			json.NewDecoder(r.Body).Decode(&params)
			log.Print(params)

			err := validateRefreshToken("TODO")
			if err != nil {
				http.Error(w, "invalid refresh token", http.StatusBadRequest)
			}
			t, _, err := createTokens("TODO")

			json.NewEncoder(w).Encode(map[string]string{
				"access_token":  t,
				"refresh_token": params.RefreshToken,
			})
		}
	}
}
