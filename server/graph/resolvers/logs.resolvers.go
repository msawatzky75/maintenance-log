package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

func (r *fuelLogResolver) ID(ctx context.Context, obj *model.FuelLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *fuelLogResolver) Vehicle(ctx context.Context, obj *model.FuelLog) (*model.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID.String())
	return &v, nil
}

func (r *maintenanceLogResolver) ID(ctx context.Context, obj *model.MaintenanceLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *maintenanceLogResolver) Vehicle(ctx context.Context, obj *model.MaintenanceLog) (*model.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID.String())
	return &v, nil
}

func (r *oilChangeLogResolver) ID(ctx context.Context, obj *model.OilChangeLog) (string, error) {
	return obj.ID.String(), nil
}

func (r *oilChangeLogResolver) Vehicle(ctx context.Context, obj *model.OilChangeLog) (*model.Vehicle, error) {
	var v model.Vehicle
	r.DB.First(&v, "id = ?", obj.VehicleID.String())
	return &v, nil
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
