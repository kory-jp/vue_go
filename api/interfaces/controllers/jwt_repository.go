package controllers

import (
	"net/http"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type JWTer interface {
	GenerateToken(ctx Context, account *domain.Account) ([]byte, error)
	GetAccountID(ctx Context) int
	GetToken(ctx Context, r *http.Request) (jwt.Token, error)
	DeleteAccountID(ctx Context, key string) error
}
