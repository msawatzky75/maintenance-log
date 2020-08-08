package graph

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenence-log/graph/model"
	uuid "github.com/satori/go.uuid"
)

// ValidateVehicle validates and returns a vehicle based on id
func ValidateVehicle(id string, db *gorm.DB) (*model.Vehicle, error) {
	var v model.Vehicle
	var c int

	i, e := uuid.FromString(id)
	if e != nil {
		return &v, fmt.Errorf("Invalid VehicleID")
	}

	db.Where("id = ?", i).First(&v).Count(&c)

	if c == 0 {
		return &v, fmt.Errorf("Unable to find vehicle")
	}

	return &v, nil
}

// ValidateUser validates and returns a user based on id
func ValidateUser(id string, db *gorm.DB) (*model.User, error) {
	var u model.User

	i, e := uuid.FromString(id)
	if e != nil {
		return &u, fmt.Errorf("Invalid UserID")
	}

	err := db.Where("id = ?", i).First(&u).Error

	if err != nil {
		return &u, fmt.Errorf("Unable to find user")
	}

	return &u, nil
}

// GetVehicleIDs returns a list of ids from an array of vehicles
func GetVehicleIDs(vs *[]*model.Vehicle) []*uuid.UUID {
	vsm := make([]*uuid.UUID, len(*vs))
	for i, v := range *vs {
		vsm[i] = &v.ID
	}
	return vsm
}
