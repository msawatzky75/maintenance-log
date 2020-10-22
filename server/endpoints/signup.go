package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

// Signup contains the options for the signup endpoints
type Signup struct {
	DB *gorm.DB
}

// SignupHandler for signup endpoint
func (s *Signup) SignupHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			s.signup(w, r)
		}
	}
}

func (s *Signup) getEmailAvailable(email string) (bool, error) {
	var u model.User
	err := s.DB.Where("email = ?", email).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		return true, nil
	}

	return false, err
}

func (s *Signup) signup(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := model.CreateUser(s.DB, creds.Email, creds.Password)
	if err == model.ErrEmailExists {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email in use."})
		return
	}
	if err == model.ErrInvalidPassword {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid password. Must be at least 8 characters."})
		return
	}
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"userId": u.ID.String()})
}
