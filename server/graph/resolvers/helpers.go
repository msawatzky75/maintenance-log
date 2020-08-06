package graph

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
	uuid "github.com/satori/go.uuid"
)

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

func ValidateUser(id string, db *gorm.DB) (*model.User, error) {
	var u model.User
	var c int

	i, e := uuid.FromString(id)
	if e != nil {
		return &u, fmt.Errorf("Invalid UserID")
	}

	db.Where("id = ?", i).First(&u).Count(&c)

	if c == 0 {
		return &u, fmt.Errorf("Unable to find user")
	}

	return &u, nil
}

func GetVehicleIDs(vs *[]*model.Vehicle) []*uuid.UUID {
	vsm := make([]*uuid.UUID, len(*vs))
	for i, v := range *vs {
		vsm[i] = &v.ID
	}
	return vsm
}
