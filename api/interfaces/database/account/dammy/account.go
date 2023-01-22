package dammy

import (
	"time"

	domain "github.com/kory-jp/vue_go/api/domain/account"
)

func AccountData() *domain.Account {
	TestAccount := &domain.Account{
		ID:        1,
		Name:      "test_user",
		Email:     "test@exm",
		Password:  "test_password",
		CreatedAt: time.Date(2023, 1, 14, 15, 55, 0, 0, time.Local),
	}
	return TestAccount
}
