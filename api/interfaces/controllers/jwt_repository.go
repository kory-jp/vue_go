package controllers

import (
	domain "github.com/kory-jp/vue_go/api/domain/account"
)

type JWTer interface {
	GenerateToken(ctx Context, account *domain.Account) ([]byte, error)
}
