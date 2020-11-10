package graph

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
	"github.com/msawatzky75/maintenance-log/server/graph/model"
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
func GetVehicleIDs(vs []*model.Vehicle) []uuid.UUID {
	vsm := make([]uuid.UUID, len(vs))
	for i, v := range vs {
		vsm[i] = v.ID
	}
	return vsm
}

// GetVehicleLogs returns logs for a vehicle
func GetVehicleLogs(DB *gorm.DB, ids []uuid.UUID, filter model.LogsFilter) ([]model.Log, error) {
	var (
		ml   []model.MaintenanceLog
		fl   []model.FuelLog
		ol   []model.OilChangeLog
		logs []model.Log
	)

	if filter.Date != nil {
		if filter.Date.EndDate.Before(filter.Date.StartDate) {
			return logs, fmt.Errorf("endDate cannot be before start date")
		}
	}

	db := DB.Where("vehicle_id in (?)", ids)
	if filter.Date != nil {
		db = db.Where("date >= ?", filter.Date.StartDate).Where("date < ?", filter.Date.EndDate)
	}
	if filter.Recent != nil && *filter.Recent > 0 {
		db = db.Order("date").Limit(*filter.Recent)
	}
	db.Find(&fl).Find(&ml).Find(&ol)

	if len(ml) > 0 {
		for _, v := range ml {
			logs = append(logs, v)
		}
	}

	if len(fl) > 0 {
		for _, v := range fl {
			logs = append(logs, v)
		}
	}

	if len(ol) > 0 {
		for _, v := range ol {
			logs = append(logs, v)
		}
	}

	if filter.Recent != nil && *filter.Recent > 0 {
		sort.Slice(logs, func(i, j int) bool {
			return logs[i].GetDate().Before(*logs[j].GetDate())
		})
		if len(logs) > *filter.Recent {
			return logs[*filter.Recent:], nil
		}
	}

	return logs, nil
}
