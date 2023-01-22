package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/kory-jp/vue_go/api/interfaces/database"
	account "github.com/kory-jp/vue_go/api/interfaces/database/account"
	usecase "github.com/kory-jp/vue_go/api/usecase/account"
)

// type AccountController struct {
// 	Interactor usecase.AccountInteractor
// }
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

func (controller *AccountController) Create(r *http.Request) (status int, message string, body interface{}, err error) {
	if r.ContentLength == 0 {
		return 400, "データ取得に失敗しました", nil, errors.New("データ取得に失敗しました")
	}
	bytesAccount, err := io.ReadAll(r.Body)
	if err != nil {
		return 400, "データ取得に失敗しました", nil, errors.New(err.Error())
	}
	accountType := new(domain.Account)
	if err := json.Unmarshal(bytesAccount, accountType); err != nil {
		return 400, "データ取得に失敗しました", nil, errors.New(err.Error())
	}
	account, err := controller.Interactor.Add(*accountType)
	if err != nil {
		return 400, err.Error(), nil, err
	}
	return 200, "新規登録完了しました", account, nil
}
