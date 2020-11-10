package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/msawatzky75/maintenance-log/server/graph/generated"
	model1 "github.com/msawatzky75/maintenance-log/server/graph/model"
)

func (r *fuelLogResolver) ID(ctx context.Context, obj *model1.FuelLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *fuelLogResolver) Vehicle(ctx context.Context, obj *model1.FuelLog) (*model1.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID)
	return &v, nil
}

func (r *fuelLogResolver) Odometer(ctx context.Context, obj *model1.FuelLog, typeArg *model1.DistanceUnit) (*model1.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

func (r *maintenanceLogResolver) ID(ctx context.Context, obj *model1.MaintenanceLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *maintenanceLogResolver) Vehicle(ctx context.Context, obj *model1.MaintenanceLog) (*model1.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID)
	return &v, nil
}

func (r *maintenanceLogResolver) Odometer(ctx context.Context, obj *model1.MaintenanceLog, typeArg *model1.DistanceUnit) (*model1.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

func (r *oilChangeLogResolver) ID(ctx context.Context, obj *model1.OilChangeLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *oilChangeLogResolver) Vehicle(ctx context.Context, obj *model1.OilChangeLog) (*model1.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID)
	return &v, nil
}

func (r *oilChangeLogResolver) Odometer(ctx context.Context, obj *model1.OilChangeLog, typeArg *model1.DistanceUnit) (*model1.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

// FuelLog returns generated1.FuelLogResolver implementation.
func (r *Resolver) FuelLog() generated1.FuelLogResolver { return &fuelLogResolver{r} }

// MaintenanceLog returns generated1.MaintenanceLogResolver implementation.
func (r *Resolver) MaintenanceLog() generated1.MaintenanceLogResolver {
	return &maintenanceLogResolver{r}
}

// OilChangeLog returns generated1.OilChangeLogResolver implementation.
func (r *Resolver) OilChangeLog() generated1.OilChangeLogResolver { return &oilChangeLogResolver{r} }

type fuelLogResolver struct{ *Resolver }
type maintenanceLogResolver struct{ *Resolver }
type oilChangeLogResolver struct{ *Resolver }
