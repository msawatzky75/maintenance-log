package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) Logs(ctx context.Context, obj *model.User, startDate int, endDate int) ([]model.Log, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Vehicles(ctx context.Context, obj *model.User) ([]*model.Vehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Preference(ctx context.Context, obj *model.User) (*model.UserPreference, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userPreferenceResolver) Vehicle(ctx context.Context, obj *model.UserPreference) (*model.Vehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserPreference returns generated.UserPreferenceResolver implementation.
func (r *Resolver) UserPreference() generated.UserPreferenceResolver {
	return &userPreferenceResolver{r}
}

type userResolver struct{ *Resolver }
type userPreferenceResolver struct{ *Resolver }
