package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) Logs(ctx context.Context, obj *model.User, startDate time.Time, endDate time.Time) ([]model.Log, error) {
	var (
		ml   []model.MaintenanceLog
		fl   []model.FuelLog
		ol   []model.OilChangeLog
		v    []*model.Vehicle
		logs []model.Log
	)

	if endDate.Before(startDate) {
		return logs, fmt.Errorf("endDate cannot be before start date")
	}

	r.DB.Find(&v, &model.Vehicle{UserID: &obj.ID})

	vIDs := GetVehicleIDs(&v)
	r.DB.Where("vehicle_id in (?)", vIDs).Where("date >= ?", startDate).Where("date < ?", endDate).Find(&fl).Find(&ml).Find(&ol)

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

	return logs, nil
}

func (r *userResolver) Vehicles(ctx context.Context, obj *model.User) ([]*model.Vehicle, error) {
	var vehicles []*model.Vehicle
	r.DB.Where(&model.Vehicle{UserID: &obj.ID}).Find(&vehicles)
	return vehicles, nil
}

func (r *userResolver) Preference(ctx context.Context, obj *model.User) (*model.UserPreference, error) {
	var p model.UserPreference
	r.DB.First(&p, "id = ?", obj.PreferenceID.String())
	return &p, nil
}

func (r *userPreferenceResolver) Vehicle(ctx context.Context, obj *model.UserPreference) (*model.Vehicle, error) {
	var v model.Vehicle
	if obj.VehicleID != nil {
		r.DB.First(&v, "id = ?", obj.VehicleID.String())
		return &v, nil
	}
	return nil, nil
}

func (r *vehicleResolver) ID(ctx context.Context, obj *model.Vehicle) (string, error) {
	return obj.ID.String(), nil
}

func (r *vehicleResolver) User(ctx context.Context, obj *model.Vehicle) (*model.User, error) {
	var u model.User
	r.DB.First(&u, "id = ?", obj.UserID.String())
	return &u, nil
}

func (r *vehicleResolver) Logs(ctx context.Context, obj *model.Vehicle, startDate time.Time, endDate time.Time) ([]model.Log, error) {
	var (
		ml   []model.MaintenanceLog
		fl   []model.FuelLog
		ol   []model.OilChangeLog
		logs []model.Log
	)

	if endDate.Before(startDate) {
		return logs, fmt.Errorf("endDate cannot be before start date")
	}

	r.DB.Where("vehicle_id = ?", obj.ID).Where("date >= ?", startDate).Where("date < ?", endDate).Find(&fl).Find(&ml).Find(&ol)

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

	return logs, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserPreference returns generated.UserPreferenceResolver implementation.
func (r *Resolver) UserPreference() generated.UserPreferenceResolver {
	return &userPreferenceResolver{r}
}

// Vehicle returns generated.VehicleResolver implementation.
func (r *Resolver) Vehicle() generated.VehicleResolver { return &vehicleResolver{r} }

type userResolver struct{ *Resolver }
type userPreferenceResolver struct{ *Resolver }
type vehicleResolver struct{ *Resolver }
