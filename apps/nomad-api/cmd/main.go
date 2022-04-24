package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nomadzetetic/apps/nomad-api/pkg/account"
	"github.com/nomadzetetic/apps/nomad-api/pkg/account/db"
	"github.com/nomadzetetic/apps/nomad-api/pkg/config"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/model"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/resolver"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/server"
	"log"
)

func graphqlHandler(graphConfig server.Config) gin.HandlerFunc {
	h := handler.NewDefaultServer(server.NewExecutableSchema(graphConfig))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextToContextMiddleware(contextKey config.GinContextKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configService := &config.EnvConfigService{}
	accountDao := db.PostgresAccountDao{Config: configService}
	accountService := account.Service{AccountDao: accountDao}
	graphConfig := server.Config{
		Resolvers: &resolver.Resolver{AccountService: accountService},
		Directives: server.DirectiveRoot{
			Authorized: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				return next(ctx)
			},
			HasRoles: func(ctx context.Context, obj interface{}, next graphql.Resolver, roles [][]model.Role) (res interface{}, err error) {
				return next(ctx)
			},
		},
	}
	port := configService.GetPort()

	r := gin.Default()
	r.Use(ginContextToContextMiddleware(configService.GetGinContextKey()))
	r.POST("/query", graphqlHandler(graphConfig))
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
