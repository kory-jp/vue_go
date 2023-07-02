package controllers

import (
	domain "github.com/kory-jp/vue_go/api/domain/account"

	"github.com/pkg/errors"

	"github.com/kory-jp/vue_go/api/interfaces/controllers"
	"github.com/kory-jp/vue_go/api/interfaces/database"
	auth "github.com/kory-jp/vue_go/api/interfaces/database/auth"
	usecase "github.com/kory-jp/vue_go/api/usecase/auth"
)

var (
	ErrMessage     = "データ取得に失敗しました"
	SuccessMessage = "ログインに成功しました"
)

type AuthController struct {
	Interactor *usecase.AuthInteractor
}

func NewAuthController(sqlHandler database.SqlHandler) *AuthController {
	return &AuthController{
		Interactor: &usecase.AuthInteractor{
			AuthRepository: &auth.AuthReporitory{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AuthController) Login(ctx controllers.Context, jwter controllers.JWTer) (status int, message string, body interface{}, err error) {
	account := new(domain.Account)
	if err = ctx.ShouldBindJSON(&account); err != nil {
		return 400, ErrMessage, nil, errors.New(err.Error())
	}
	if err = controller.Interactor.Auth(*account); err != nil {
		return 400, ErrMessage, nil, errors.New(err.Error())
	}
	jwt, err := jwter.GenerateToken(ctx, account)
	if err != nil {
		return 400, ErrMessage, nil, errors.New(err.Error())
	}
	return 200, SuccessMessage, jwt, nil
}

// func (controller *AuthController) Logout(ctx controllers.Context, jwter controllers.JWTer) (status int, message string, body interface{}, err error) {
// 	accountId, ok := jwter.GetAccountID(ctx)
// 	if !ok {
// 		return 400, ErrMessage, nil, nil
// 	}
// 	aid := accountId.(domain.Account.ID)

// 	if err :=
// }
