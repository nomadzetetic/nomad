package account

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/nomadzetetic/apps/nomad-api/pkg/config"
	"github.com/nomadzetetic/apps/nomad-api/pkg/graph/model"
	"github.com/nomadzetetic/apps/nomad-api/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	AccountDao    Dao
	ConfigService config.Service
}

func NewAccountService(accountDao Dao, configService config.Service) *Service {
	return &Service{
		AccountDao:    accountDao,
		ConfigService: configService,
	}
}

func (s Service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s Service) Register(ctx context.Context, input model.RegisterAccountInput) (bool, error) {
	obfuscatedPassword, hashError := s.hashPassword(input.Password)
	if hashError != nil {
		return false, hashError
	}

	_, dbError := s.AccountDao.Create(ctx, CreateAccountParams{
		Email:              input.Email,
		Nickname:           input.Nickname,
		ObfuscatedPassword: obfuscatedPassword,
		ActivationToken:    uuid.New().String(),
	})

	if dbError != nil {
		return false, dbError
	}

	return true, nil
}

func (s Service) Login(ctx context.Context, input model.LoginInput) (*model.LoginResult, error) {
	account, err := s.AccountDao.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if account.ActivationToken != nil {
		return nil, fmt.Errorf("account not active")
	}

	if account.Banned == true {
		return nil, fmt.Errorf("account banned")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.ObfuscatedPassword), []byte(input.Password))

	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	token, err := jwt.Sign([]byte(s.ConfigService.GetJwtSecret()), account.ID, account.Roles, nil)
	if err != nil {
		return nil, err
	}

	return &model.LoginResult{
		JwtToken: *token,
	}, nil
}

func (s Service) Activate(ctx context.Context, input model.ActivateAccountInput) (bool, error) {
	account, err := s.AccountDao.Activate(ctx, ActivateAccountParams{
		ID:              input.AccountID,
		ActivationToken: input.ActivationToken,
	})

	if err != nil {
		return false, err
	}

	if account != nil {
		return true, nil
	}

	return false, nil
}
