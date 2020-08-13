package endpoints

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func Refresh(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("refresh hit")
	}
}
