package usecase

import (
	domain "github.com/kory-jp/vue_go/api/domain/account"
)

type AccountRepository interface {
	Store(domain.Account) (int, error)
	FindById(int) (*domain.Account, error)
	FindByEmail(string) (int, error)
}
