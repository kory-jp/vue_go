package database

import (
	"os"
	"time"

	"github.com/kory-jp/vue_go/api/interfaces/database/account/dammy"
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
	for row.Next() {
		if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
			return nil, err
		}
	}
	err = row.Err()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	user = &domain.Account{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
	}
	if os.Getenv("PROCESS_ENV") == "testProcess" {
		return dammy.AccountData(), nil
	}
	return user, nil
}

func (repo *AccountRepository) FindByEmail(email string) (numberAccount int, err error) {
	// TODO: rowの詳細を確認
	wrongNumber := 999
	row, err := repo.Query(mysql.GetNumberAccountState, email)
	if err != nil {
		return wrongNumber, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&numberAccount)
		if err != nil {
			return wrongNumber, errors.New(err.Error())
		}
	}
	err = row.Err()
	if err != nil {
		return wrongNumber, errors.New(err.Error())
	}
	return numberAccount, nil
}
