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
	ErrMessage           = "データ取得に失敗しました"
	SuccessLoginMessage  = "ログインに成功しました"
	SuccessLogoutMessage = "ログアウトに成功しました"
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

func (controller *AuthController) Login(ctx controllers.Context, jwter controllers.JWTer) (status int, message string, err error) {
	account := new(domain.Account)
	if err = ctx.ShouldBindJSON(&account); err != nil {
		return 400, ErrMessage, errors.New(err.Error())
	}
	account, err = controller.Interactor.Auth(*account)
	if err != nil {
		return 400, ErrMessage, errors.New(err.Error())
	}
	jwt, err := jwter.GenerateToken(ctx, account)
	if err != nil {
		return 400, ErrMessage, errors.New(err.Error())
	}
	ctx.SetCookie("token", string(jwt), 60, "/", "http://localhost", false, true)
	return 200, SuccessLoginMessage, nil
}

func (controller *AuthController) Logout(ctx controllers.Context, jwter controllers.JWTer, tokenID string) (status int, message string, err error) {
	if err = jwter.DeleteAccountID(ctx, tokenID); err != nil {
		return 400, ErrMessage, err
	}
	return 200, SuccessLogoutMessage, nil
}
