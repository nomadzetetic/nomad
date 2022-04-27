package resolver

import "github.com/nomadzetetic/apps/nomad-api/pkg/account"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AccountService *account.Service
}
