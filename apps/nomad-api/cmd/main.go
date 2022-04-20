package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nomadzetetic/apps/nomad-api/pkg/config"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/resolvers"
)

func graphqlHandler(gqlServerConfig graph.Config) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(gqlServerConfig))

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
	configService := &config.EnvConfigService{}
	port := configService.GetPort()
	graphConfig := graph.Config{Resolvers: &resolvers.Resolver{}}
	graph.SetupDirectives(graphConfig)

	r := gin.Default()
	r.Use(ginContextToContextMiddleware(configService.GetGinContextKey()))
	r.POST("/query", graphqlHandler(graphConfig))
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
