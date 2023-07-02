package controllers

import (
	"github.com/kory-jp/vue_go/api/interfaces/controllers"

	"github.com/pkg/errors"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/kory-jp/vue_go/api/interfaces/database"
	account "github.com/kory-jp/vue_go/api/interfaces/database/account"
	usecase "github.com/kory-jp/vue_go/api/usecase/account"
)

type AccountController struct {
	Interactor AccountInteractor
}

func NewAccountController(sqlHandler database.SqlHandler) *AccountController {
	return &AccountController{
		Interactor: &usecase.AccountInteractor{
			AccountRepository: &account.AccountRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AccountController) Create(ctx controllers.Context) (status int, message string, body interface{}, err error) {
	accountType := new(domain.Account)
	if err = ctx.ShouldBindJSON(&accountType); err != nil {
		return 400, "データ取得に失敗しました", nil, errors.New(err.Error())
	}
	account, err := controller.Interactor.Add(*accountType)
	if err != nil {
		return 400, err.Error(), nil, err
	}
	return 200, "新規登録完了しました", account, nil
}
