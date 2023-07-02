package usecase

import (
	domain "github.com/kory-jp/vue_go/api/domain/account"
	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor struct {
	AuthRepository AuthRepository
}

func (interactor *AuthInteractor) Auth(ac domain.Account) error {
	account, err := interactor.AuthRepository.GetByEmail(ac.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(ac.Password))
	if err != nil {
		return err
	}
	return nil
}
