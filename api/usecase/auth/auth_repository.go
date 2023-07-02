package usecase

import (
	domain "github.com/kory-jp/vue_go/api/domain/account"
)

type AuthRepository interface {
	GetByEmail(string) (*domain.Account, error)
}
