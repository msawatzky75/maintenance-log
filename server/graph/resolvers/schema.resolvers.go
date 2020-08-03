package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
	uuid "github.com/satori/go.uuid"
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
	userID, err := uuid.FromString(data.UserID)
	if err != nil {
		panic("Invalid UserID")
	}

	v := &model.Vehicle{
		Make:     data.Make,
		Model:    data.Model,
		Year:     data.Year,
		Odometer: &model.DistanceValue{Value: &data.Odometer.Value, Type: &data.Odometer.Type},
		UserID:   &userID,
	}

	r.DB.Create(v)
	return v, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id *string) (*model.User, error) {
	uid, err := uuid.FromString(*id)
	if err != nil {
		panic("Unable to convert string to uuid")
	}

	u := model.User{}
	r.DB.First(&u, "id = ?", uid)

	return &u, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
