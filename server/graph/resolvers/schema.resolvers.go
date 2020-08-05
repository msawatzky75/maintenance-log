package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserInput) (*model.User, error) {
	tx := r.DB.Begin()
	// Rollback the transaction if the routine panics
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	p := &model.UserPreference{}
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	u := &model.User{Email: data.Email, PreferenceID: &p.ID}
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return u, tx.Commit().Error
}

func (r *mutationResolver) CreateVehicle(ctx context.Context, data model.VehicleInput) (*model.Vehicle, error) {
	u, err := ValidateUser(data.UserID, r.DB)
	if err != nil {
		return &model.Vehicle{}, err
	}

	fmt.Println(u.ID)

	v := &model.Vehicle{
		Make:     data.Make,
		Model:    data.Model,
		Year:     data.Year,
		Odometer: &model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type},
		UserID:   &u.ID,
	}

	r.DB.Create(v)
	return v, nil
}

func (r *mutationResolver) CreateFuelLog(ctx context.Context, data model.FuelLogInput) (*model.FuelLog, error) {
	v, err := ValidateVehicle(data.VehicleID, r.DB)
	if err != nil {
		return &model.FuelLog{}, err
	}
	odometer := model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type}
	trip := model.DistanceValue{Value: &data.Trip.Value, Type: &data.Trip.Type}
	fuelAmount := model.FluidValue{Value: &data.FuelAmount.Value, Type: &data.FuelAmount.Type}
	fuelPrice := model.MoneyValue{Value: &data.FuelPrice.Value, Type: &data.FuelPrice.Type}
	l := model.FuelLog{
		Date:       &data.Date,
		VehicleID:  &v.ID,
		Odometer:   &odometer,
		Notes:      data.Notes,
		Trip:       &trip,
		Grade:      &data.Grade,
		FuelAmount: &fuelAmount,
		FuelPrice:  &fuelPrice,
	}
	return &l, nil
}

func (r *mutationResolver) CreateMaintenanceLog(ctx context.Context, data model.MaintenanceLogInput) (*model.MaintenanceLog, error) {
	v, err := ValidateVehicle(data.VehicleID, r.DB)
	if err != nil {
		return &model.MaintenanceLog{}, err
	}
	o := model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type}
	l := model.MaintenanceLog{Date: &data.Date, VehicleID: &v.ID, Odometer: &o, Notes: data.Notes}
	return &l, nil
}

func (r *mutationResolver) CreateOilChangeLog(ctx context.Context, data model.OilChangeLogInput) (*model.OilChangeLog, error) {
	v, err := ValidateVehicle(data.VehicleID, r.DB)
	if err != nil {
		return &model.OilChangeLog{}, err
	}
	o := model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type}
	l := model.OilChangeLog{Date: &data.Date, VehicleID: &v.ID, Odometer: &o, Notes: data.Notes}
	return &l, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id *string) (*model.User, error) {
	u, err := ValidateUser(*id, r.DB)
	if err != nil {
		return &model.User{}, err
	}

	return u, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
