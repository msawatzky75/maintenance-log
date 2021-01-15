package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/msawatzky75/maintenance-log/server/graph/generated"
	"github.com/msawatzky75/maintenance-log/server/graph/model"
)

func (r *fuelLogResolver) ID(ctx context.Context, obj *model.FuelLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *fuelLogResolver) Vehicle(ctx context.Context, obj *model.FuelLog) (*model.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID)
	return &v, nil
}

func (r *fuelLogResolver) Odometer(ctx context.Context, obj *model.FuelLog, typeArg *model.DistanceUnit) (*model.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

func (r *fuelLogResolver) Efficiency(ctx context.Context, obj *model.FuelLog) (*model.FuelEfficiency, error) {
	litre := obj.FuelAmount.ConvertTo(model.FluidUnitLitre).Value
	gallon := obj.FuelAmount.ConvertTo(model.FluidUnitGallon).Value
	km := obj.Trip.ConvertTo(model.DistanceUnitKilometre).Value
	mile := obj.Trip.ConvertTo(model.DistanceUnitMile).Value

	e := model.FuelEfficiency{
		Kml:   *km / *litre,
		L100k: (100 * *litre) / *km,
		Mpg:   *mile / *gallon,
	}

	return &e, nil
}

func (r *maintenanceLogResolver) ID(ctx context.Context, obj *model.MaintenanceLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *maintenanceLogResolver) Vehicle(ctx context.Context, obj *model.MaintenanceLog) (*model.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID)
	return &v, nil
}

func (r *maintenanceLogResolver) Odometer(ctx context.Context, obj *model.MaintenanceLog, typeArg *model.DistanceUnit) (*model.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

func (r *oilChangeLogResolver) ID(ctx context.Context, obj *model.OilChangeLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *oilChangeLogResolver) Vehicle(ctx context.Context, obj *model.OilChangeLog) (*model.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID)
	return &v, nil
}

func (r *oilChangeLogResolver) Odometer(ctx context.Context, obj *model.OilChangeLog, typeArg *model.DistanceUnit) (*model.DistanceValue, error) {
	if typeArg != nil {
		return obj.Odometer.ConvertTo(*typeArg), nil
	}
	return obj.Odometer, nil
}

// FuelLog returns generated.FuelLogResolver implementation.
func (r *Resolver) FuelLog() generated.FuelLogResolver { return &fuelLogResolver{r} }

// MaintenanceLog returns generated.MaintenanceLogResolver implementation.
func (r *Resolver) MaintenanceLog() generated.MaintenanceLogResolver {
	return &maintenanceLogResolver{r}
}

// OilChangeLog returns generated.OilChangeLogResolver implementation.
func (r *Resolver) OilChangeLog() generated.OilChangeLogResolver { return &oilChangeLogResolver{r} }

type fuelLogResolver struct{ *Resolver }
type maintenanceLogResolver struct{ *Resolver }
type oilChangeLogResolver struct{ *Resolver }
