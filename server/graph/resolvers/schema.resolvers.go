package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
	"github.com/msawatzky75/maintenence-log/server/middleware"
	uuid "github.com/satori/go.uuid"
)

func (r *mutationResolver) CreateVehicle(ctx context.Context, data model.VehicleInput) (*model.Vehicle, error) {
	uid, err := uuid.FromString((*ctx.Value(middleware.JwtContextKey).(*jwt.MapClaims))["userId"].(string))
	if err != nil {
		return nil, err
	}

	v := &model.Vehicle{
		Make:     data.Make,
		Model:    data.Model,
		Year:     data.Year,
		Odometer: &model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type},
		UserID:   &uid,
	}

	r.DB.Create(v)
	return v, nil
}

func (r *mutationResolver) CreateFuelLog(ctx context.Context, data model.FuelLogInput) (*model.FuelLog, error) {
	tx := r.DB.Begin()
	// Rollback the transaction if the routine panics
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	v, err := ValidateVehicle(data.VehicleID, r.DB)
	if err != nil {
		tx.Rollback()
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

	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	v.Odometer = l.Odometer

	if err := tx.Save(&v).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &l, tx.Commit().Error
}

func (r *mutationResolver) CreateMaintenanceLog(ctx context.Context, data model.MaintenanceLogInput) (*model.MaintenanceLog, error) {
	tx := r.DB.Begin()
	// Rollback the transaction if the routine panics
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	v, err := ValidateVehicle(data.VehicleID, r.DB)
	if err != nil {
		tx.Rollback()
		return &model.MaintenanceLog{}, err
	}

	o := model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type}
	l := model.MaintenanceLog{Date: &data.Date, VehicleID: &v.ID, Odometer: &o, Notes: data.Notes}

	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	v.Odometer = l.Odometer

	if err := tx.Save(&v).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &l, tx.Commit().Error
}

func (r *mutationResolver) CreateOilChangeLog(ctx context.Context, data model.OilChangeLogInput) (*model.OilChangeLog, error) {
	tx := r.DB.Begin()
	// Rollback the transaction if the routine panics
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	v, err := ValidateVehicle(data.VehicleID, r.DB)
	if err != nil {
		tx.Rollback()
		return &model.OilChangeLog{}, err
	}

	o := model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type}
	l := model.OilChangeLog{Date: &data.Date, VehicleID: &v.ID, Odometer: &o, Notes: data.Notes}

	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	v.Odometer = l.Odometer

	if err := tx.Save(&v).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &l, tx.Commit().Error
}

func (r *mutationResolver) UpdatePreference(ctx context.Context, data model.UserPreferenceInput) (*model.UserPreference, error) {
	var (
		p   model.UserPreference
		u   *model.User
		err error
	)
	uid := (*ctx.Value(middleware.JwtContextKey).(*jwt.MapClaims))["userId"].(string)

	u, err = ValidateUser(uid, r.DB)
	if err != nil {
		return &p, err
	}

	if u.PreferenceID != nil {
		r.DB.Where("id = ?", u.PreferenceID).Find(&p)
	}

	if data.VehicleID != nil {
		v, err := ValidateVehicle(*data.VehicleID, r.DB)
		if err != nil {
			return &p, err
		}
		p.VehicleID = &v.ID
	}

	if data.Distance != nil {
		p.Distance = data.Distance
	}

	if data.Fluid != nil {
		p.Fluid = data.Fluid
	}

	if data.Money != nil {
		p.Money = data.Money
	}

	r.DB.Save(&p)

	return &p, nil
}

func (r *mutationResolver) DeleteFuelLog(ctx context.Context, id string) (*bool, error) {
	var (
		err error
		s   = true
	)

	err = r.DB.Delete(&model.FuelLog{}).Where("id = ?", id).Error
	if err != nil {
		fmt.Print(err)
		s = false
	}

	return &s, err
}

func (r *mutationResolver) DeleteMaintenanceLog(ctx context.Context, id string) (*bool, error) {
	var (
		err error
		s   = true
	)

	err = r.DB.Delete(&model.MaintenanceLog{}).Where("id = ?", id).Error
	if err != nil {
		fmt.Print(err)
		s = false
	}

	return &s, err
}

func (r *mutationResolver) DeleteOilChangeLog(ctx context.Context, id string) (*bool, error) {
	var (
		err error
		s   = true
	)

	err = r.DB.Delete(&model.OilChangeLog{}).Where("id = ?", id).Error
	if err != nil {
		fmt.Print(err)
		s = false
	}

	return &s, err
}

func (r *queryResolver) GetUser(ctx context.Context, id *string) (*model.User, error) {
	var (
		u   *model.User
		err error
	)
	uid := (*ctx.Value(middleware.JwtContextKey).(*jwt.MapClaims))["userId"].(string)

	if id != nil {
		u, err = ValidateUser(*id, r.DB)
		if err != nil {
			return &model.User{}, err
		}
	} else {
		u, err = ValidateUser(uid, r.DB)
		if err != nil {
			return &model.User{}, err
		}
	}

	return u, nil
}

func (r *queryResolver) GetVehicle(ctx context.Context, id string) (*model.Vehicle, error) {
	var v model.Vehicle
	uid, err := uuid.FromString(id)
	if err != nil {
		return &v, err
	}
	r.DB.Where("id = ?", uid).Find(&v)
	return &v, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
