package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nomadzetetic/apps/nomad-api/pkg/graph"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/models"
)

func (r *mutationResolver) RegisterAccount(ctx context.Context, input models.RegisterAccountInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ActivateAccount(ctx context.Context, input models.ActivateAccountInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Account(ctx context.Context) (*models.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, input models.LoginInput) (*models.LoginResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
