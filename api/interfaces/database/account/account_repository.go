package database

import (
	"time"

	"github.com/pkg/errors"

	"github.com/kory-jp/vue_go/api/interfaces/database/account/mysql"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/kory-jp/vue_go/api/interfaces/database"
)

type AccountRepository struct {
	database.SqlHandler
}

func (repo *AccountRepository) Store(ac domain.Account) (id int, err error) {
	result, err := repo.Execute(mysql.CreateAccountState, ac.Name, ac.Email, ac.Password)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New(err.Error())
	}
	id = int(id64)
	return id, nil
}

func (repo *AccountRepository) FindById(identifier int) (user *domain.Account, err error) {
	row, err := repo.Query(mysql.FindAccountState, identifier)
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
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
		return nil, err
	}
	user = &domain.Account{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
	}
	return user, nil
}

func (repo *AccountRepository) FindByEmail(email string) (numberAccount int, err error) {
	// TODO: rowの詳細を確認
	row, err := repo.Query(mysql.GetNumberAccountState, email)
	if err != nil {
		return 0, err
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&numberAccount)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return numberAccount, nil
}
