package account

import (
	"github.com/kory-jp/vue_go/api/domain"
)

type AccountRepository interface {
	Store(domain.Account) (int, error)
	FindById(int) (*domain.Account, error)
}
