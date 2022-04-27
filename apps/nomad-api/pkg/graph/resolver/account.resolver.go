package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/model"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/server"
)

func (r *mutationResolver) RegisterAccount(ctx context.Context, input model.RegisterAccountInput) (bool, error) {
	return r.AccountService.Register(ctx, input)
}

func (r *mutationResolver) ActivateAccount(ctx context.Context, input model.ActivateAccountInput) (bool, error) {
	return r.AccountService.Activate(ctx, input)
}

func (r *queryResolver) Account(ctx context.Context) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
