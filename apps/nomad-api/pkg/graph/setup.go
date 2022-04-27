package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nomadzetetic/apps/nomad-api/pkg/account"
	"github.com/nomadzetetic/apps/nomad-api/pkg/config"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/model"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/resolver"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/server"
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextMiddleware(contextKey config.GinContextKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func graphqlHandler(resolver resolver.Resolver) gin.HandlerFunc {
	graphConfig := server.Config{
		Resolvers: &resolver,
		Directives: server.DirectiveRoot{
			Authorized: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				return next(ctx)
			},
			HasRoles: func(ctx context.Context, obj interface{}, next graphql.Resolver, roles [][]model.Role) (res interface{}, err error) {
				return next(ctx)
			},
		},
	}

	h := handler.NewDefaultServer(server.NewExecutableSchema(graphConfig))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetupGinEndpoints(configService config.Service, router *gin.Engine, accountService *account.Service) {
	r := resolver.Resolver{AccountService: accountService}
	router.Use(ginContextMiddleware(configService.GetGinContextKey()))
	router.POST("/query", graphqlHandler(r))
	router.GET("/", playgroundHandler())
}
