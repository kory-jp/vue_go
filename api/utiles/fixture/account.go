package fixture

import (
	"math/rand"
	"strconv"
	"time"

	domain "github.com/kory-jp/vue_go/api/domain/account"
)

func Account(ac *domain.Account) *domain.Account {
	result := &domain.Account{
		ID:        rand.Int(),
		Name:      "sampleTest" + strconv.Itoa(rand.Int())[:5],
		Email:     "sample@exm.com",
		Password:  "password",
		CreatedAt: time.Now(),
	}
	if ac == nil {
		return result
	}

	if ac.ID != 0 {
		result.ID = ac.ID
	}

	if ac.Name != "" {
		result.Name = ac.Name
	}

	if ac.Email != "" {
		result.Email = ac.Email
	}

	if ac.Password != "" {
		result.Password = ac.Password
	}

	if !ac.CreatedAt.IsZero() {
		result.CreatedAt = ac.CreatedAt
	}

	return result
}
