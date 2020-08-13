package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

// User is the user endpoint handler
func User(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			u, err := model.CreateUser(db, r.FormValue("email"), r.FormValue("password"))
			if err != nil {
				http.Error(w, err.Error(), 500)

			}

			// Log them in
			t, rt, err := createTokens(u.ID.String())

			json.NewEncoder(w).Encode(map[string]string{
				"access_token":  t,
				"refresh_token": rt,
			})
		default:
			fmt.Fprint(w, "Only POST is supported for this endpoint")
		}
	}
}
