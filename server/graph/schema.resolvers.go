package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"

	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserInput) (*model.User, error) {
	u := model.User{Email: data.Email}
	r.DB.Create(&u)
	return &u, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id *string) (*model.User, error) {
	uid, err := uuid.FromString(*id)
	if err != nil {
		panic("Unable to convert string to uuid")
	}

	log.Print(r.DB.Where("id = ?", uid).First(&model.User{}))

	return &model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
