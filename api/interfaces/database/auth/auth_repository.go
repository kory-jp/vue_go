package database

import (
	"time"

	"github.com/pkg/errors"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/kory-jp/vue_go/api/interfaces/database"
	"github.com/kory-jp/vue_go/api/interfaces/database/auth/mysql"
)

type AuthReporitory struct {
	database.SqlHandler
}

func (repo *AuthReporitory) GetByEmail(getEmail string) (*domain.Account, error) {
	row, err := repo.Query(mysql.GetAccountByEmail, getEmail)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var (
		id         int
		name       string
		email      string
		password   string
		created_at time.Time
	)
	for row.Next() {
		if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
			return nil, err
		}
	}
	err = row.Err()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	account := &domain.Account{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
	}
	return account, nil
}
