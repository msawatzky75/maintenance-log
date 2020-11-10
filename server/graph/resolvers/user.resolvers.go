package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/msawatzky75/maintenance-log/server/graph/generated"
	model1 "github.com/msawatzky75/maintenance-log/server/graph/model"
	"github.com/satori/go.uuid"
)

func (r *userResolver) ID(ctx context.Context, obj *model1.User) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) Logs(ctx context.Context, obj *model1.User, filter model1.LogsFilter) ([]model1.Log, error) {
	var v []*model.Vehicle

	r.DB.Find(&v, &model.Vehicle{UserID: &obj.ID})

	return GetVehicleLogs(r.DB, GetVehicleIDs(v), filter)
}

func (r *userResolver) Vehicles(ctx context.Context, obj *model1.User) ([]*model1.Vehicle, error) {
	var vehicles []*model.Vehicle
	r.DB.Where(&model.Vehicle{UserID: &obj.ID}).Find(&vehicles)
	return vehicles, nil
}

func (r *userResolver) Preference(ctx context.Context, obj *model1.User) (*model1.UserPreference, error) {
	var p model.UserPreference
	r.DB.First(&p, "id = ?", obj.PreferenceID.String())
	return &p, nil
}

func (r *userPreferenceResolver) Vehicle(ctx context.Context, obj *model1.UserPreference) (*model1.Vehicle, error) {
	var v model.Vehicle
	if obj.VehicleID != nil {
		r.DB.First(&v, "id = ?", obj.VehicleID.String())
		return &v, nil
	}
	return nil, nil
}

func (r *vehicleResolver) ID(ctx context.Context, obj *model1.Vehicle) (string, error) {
	return obj.ID.String(), nil
}

func (r *vehicleResolver) User(ctx context.Context, obj *model1.Vehicle) (*model1.User, error) {
	var u model.User
	r.DB.First(&u, "id = ?", obj.UserID.String())
	return &u, nil
}

func (r *vehicleResolver) Odometer(ctx context.Context, obj *model1.Vehicle, typeArg *model1.DistanceUnit) (*model1.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

func (r *vehicleResolver) Logs(ctx context.Context, obj *model1.Vehicle, filter model1.LogsFilter) ([]model1.Log, error) {
	return GetVehicleLogs(r.DB, []uuid.UUID{obj.ID}, filter)
}

// User returns generated1.UserResolver implementation.
func (r *Resolver) User() generated1.UserResolver { return &userResolver{r} }

// UserPreference returns generated1.UserPreferenceResolver implementation.
func (r *Resolver) UserPreference() generated1.UserPreferenceResolver {
	return &userPreferenceResolver{r}
}

// Vehicle returns generated1.VehicleResolver implementation.
func (r *Resolver) Vehicle() generated1.VehicleResolver { return &vehicleResolver{r} }

type userResolver struct{ *Resolver }
type userPreferenceResolver struct{ *Resolver }
type vehicleResolver struct{ *Resolver }
