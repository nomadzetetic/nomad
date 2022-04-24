package db

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nomadzetetic/apps/nomad-api/pkg/config"
)

type AccountDao interface {
	Create(ctx context.Context, params CreateAccountParams) (*AccountEntity, error)
	GetByEmail(ctx context.Context, email string) (*AccountEntity, error)
}

type PostgresAccountDao struct {
	Config config.Service
}

type CreateAccountParams struct {
	Email              string
	Nickname           string
	ObfuscatedPassword string
}

func (accountDao PostgresAccountDao) Create(ctx context.Context, params CreateAccountParams) (*AccountEntity, error) {
	db, connErr := pgxpool.Connect(ctx, accountDao.Config.GetPostgresDatabaseUrl())
	if connErr != nil {
		return nil, connErr
	}
	defer db.Close()

	var id string
	insertErr := db.QueryRow(
		context.Background(),
		"insert into accounts(email,nickname,obfuscated_password,enabled) values ($1,$2,$3,true) returning id",
		params.Email,
		params.Nickname,
		params.ObfuscatedPassword,
	).Scan(&id)

	if insertErr != nil {
		return nil, insertErr
	}

	var accounts []*AccountEntity
	selectError := pgxscan.Select(ctx, db, &accounts, "select * from accounts where id = $1", id)

	if selectError != nil {
		return nil, selectError
	}

	if len(accounts) != 1 {
		return nil, fmt.Errorf("register account reselect error id=\"%s\"", id)
	}

	return accounts[0], nil
}

func (accountDao PostgresAccountDao) GetByEmail(ctx context.Context, email string) (*AccountEntity, error) {
	db, connErr := pgxpool.Connect(ctx, accountDao.Config.GetPostgresDatabaseUrl())
	if connErr != nil {
		return nil, connErr
	}
	defer db.Close()

	var accounts []*AccountEntity
	selectError := pgxscan.Select(ctx, db, &accounts, "select * from accounts where email = $1", email)

	if selectError != nil {
		return nil, selectError
	}

	if len(accounts) != 1 {
		return nil, fmt.Errorf("invalid email or password")
	}

	return accounts[0], nil
}
