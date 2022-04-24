package account

import (
	"context"
	"fmt"
	"github.com/nomadzetetic/apps/nomad-api/pkg/account/db"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/model"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	AccountDao db.AccountDao
}

func (accountService Service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (accountService Service) RegisterAccount(ctx context.Context, input model.RegisterAccountInput) (bool, error) {
	obfuscatedPassword, hashError := accountService.HashPassword(input.Password)
	if hashError != nil {
		return false, hashError
	}

	_, dbError := accountService.AccountDao.Create(ctx, db.CreateAccountParams{
		Email:              input.Email,
		Nickname:           input.Nickname,
		ObfuscatedPassword: obfuscatedPassword,
	})

	if dbError != nil {
		return false, dbError
	}

	return true, nil
}

func (accountService Service) Login(ctx context.Context, input model.LoginInput) (*model.LoginResult, error) {
	account, getByEmailErr := accountService.AccountDao.GetByEmail(ctx, input.Email)
	if getByEmailErr != nil {
		return nil, getByEmailErr
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(account.ObfuscatedPassword), []byte(input.Password))

	if compareErr != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return &model.LoginResult{
		// TODO: Implement JWT service
		JwtToken: "",
	}, nil
}
