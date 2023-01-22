package controllers

import (
	"context"

	domain "github.com/kory-jp/vue_go/api/domain/account"
)

type KVS interface {
	Save(context.Context, string, domain.Account) error
	Load(context.Context, string) (*domain.Account, error)
}
