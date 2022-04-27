package account

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dao interface {
	GetById(ctx context.Context, id string) (*Entity, error)
	GetByEmail(ctx context.Context, email string) (*Entity, error)
	Create(ctx context.Context, params CreateAccountParams) (*Entity, error)
	Activate(ctx context.Context, params ActivateAccountParams) (*Entity, error)
}

type CreateAccountParams struct {
	Email              string
	Nickname           string
	ObfuscatedPassword string
	ActivationToken    string
}

type ActivateAccountParams struct {
	ID              string
	ActivationToken string
}

type PostgresDao struct {
	pool *pgxpool.Pool
}

func NewAccountDao(pool *pgxpool.Pool) *PostgresDao {
	return &PostgresDao{
		pool: pool,
	}
}

func (dao PostgresDao) Create(ctx context.Context, params CreateAccountParams) (*Entity, error) {
	var id string
	insertErr := dao.pool.QueryRow(
		ctx,
		"insert into accounts(email,nickname,obfuscated_password,activation_token,banned) values ($1,$2,$3,$4,false) returning id",
		params.Email,
		params.Nickname,
		params.ObfuscatedPassword,
		params.ActivationToken,
	).Scan(&id)

	if insertErr != nil {
		return nil, insertErr
	}

	var accounts []*Entity
	selectError := pgxscan.Select(ctx, dao.pool, &accounts, "select * from accounts where id = $1", id)

	if selectError != nil {
		return nil, selectError
	}

	if len(accounts) != 1 {
		return nil, fmt.Errorf("register account reselect error id=\"%s\"", id)
	}

	return accounts[0], nil
}

func (dao PostgresDao) GetById(ctx context.Context, id string) (*Entity, error) {
	var accounts []*Entity
	selectError := pgxscan.Select(ctx, dao.pool, &accounts, "select * from accounts where id = $1", id)

	if selectError != nil {
		return nil, selectError
	}

	if len(accounts) == 0 {
		return nil, nil
	}

	return accounts[0], nil
}

func (dao PostgresDao) GetByEmail(ctx context.Context, email string) (*Entity, error) {
	var accounts []*Entity
	selectError := pgxscan.Select(ctx, dao.pool, &accounts, "select * from accounts where email = $1", email)

	if selectError != nil {
		return nil, selectError
	}

	if len(accounts) == 0 {
		return nil, nil
	}

	return accounts[0], nil
}

func (dao PostgresDao) Activate(ctx context.Context, params ActivateAccountParams) (*Entity, error) {
	var id string
	updateErr := dao.pool.QueryRow(
		ctx,
		"update accounts set activation_token = null where id = $1 and activation_token = $2 returning id",
		params.ID,
		params.ActivationToken,
	).Scan(&id)

	if updateErr != nil {
		if updateErr.Error() == "no rows in result set" {
			return nil, fmt.Errorf("invalid activation token")
		}
		return nil, updateErr
	}

	return dao.GetById(ctx, params.ID)
}
