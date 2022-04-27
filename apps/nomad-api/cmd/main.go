package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nomadzetetic/apps/nomad-api/pkg/account"
	"github.com/nomadzetetic/apps/nomad-api/pkg/config"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph"
	"log"
)

func main() {
	configService := config.NewConfig()

	pool, err := pgxpool.Connect(context.Background(), configService.GetPostgresDatabaseUrl())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	accountDao := account.NewAccountDao(pool)
	accountService := account.NewAccountService(accountDao, configService)

	router := gin.Default()

	graph.SetupGinEndpoints(configService, router, accountService)

	err = router.Run(":" + configService.GetPort())
	if err != nil {
		log.Fatal(err)
	}
}
