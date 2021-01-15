package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/msawatzky75/maintenance-log/server/graph/generated"
	"github.com/msawatzky75/maintenance-log/server/graph/model"
	"github.com/satori/go.uuid"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) Logs(ctx context.Context, obj *model.User, filter model.LogsFilter) ([]model.Log, error) {
	var v []*model.Vehicle

	r.DB.Find(&v, &model.Vehicle{UserID: &obj.ID})

	return GetVehicleLogs(r.DB, GetVehicleIDs(v), filter)
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

func (r *vehicleResolver) Odometer(ctx context.Context, obj *model.Vehicle, typeArg *model.DistanceUnit) (*model.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

func (r *vehicleResolver) Logs(ctx context.Context, obj *model.Vehicle, filter model.LogsFilter) ([]model.Log, error) {
	return GetVehicleLogs(r.DB, []uuid.UUID{obj.ID}, filter)
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
