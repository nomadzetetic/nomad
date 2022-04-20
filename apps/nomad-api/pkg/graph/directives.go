package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/models"
)

func SetupDirectives(config Config) {
	config.Directives.Authorized = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		return next(ctx)
	}

	config.Directives.HasRoles = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles [][]models.Role) (res interface{}, err error) {
		return next(ctx)
	}
}
