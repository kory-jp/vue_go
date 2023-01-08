package database

import (
	"fmt"
	"log"
	"time"

	"github.com/kory-jp/vue_go/api/interfaces/database/account/mysql"

	"github.com/kory-jp/vue_go/api/domain"
	"github.com/kory-jp/vue_go/api/interfaces/database"
)

type AccountRepository struct {
	database.SqlHandler
}

func (repo *AccountRepository) Store(ac domain.Account) (id int, err error) {
	result, err := repo.Execute(mysql.CreateAccountState, ac.Name, ac.Email, ac.Password)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	id = int(id64)
	return id, nil
}

func (repo *AccountRepository) FindById(identifier int) (user *domain.Account, err error) {
	row, err := repo.Query(mysql.FindAccountState, identifier)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
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
		fmt.Println(err)
		log.Println(err)
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
