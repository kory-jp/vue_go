package usecase

import (
	domain "github.com/kory-jp/vue_go/api/domain/account"
)

type AccountInteractor struct {
	AccountRepository AccountRepository
}

func (interactor *AccountInteractor) Add(ac domain.Account) (account *domain.Account, err error) {
	if err = ac.AccountValidate(); err == nil {
		if err = ac.CheckExistEmail(interactor.AccountRepository.FindByEmail(ac.Email)); err == nil {
			ac.Password, err = ac.Encrypt(ac.Password)
			if err != nil {
				return nil, err
			}
			identifier, err := interactor.AccountRepository.Store(ac)
			if err != nil {
				return nil, err
			} else {
				account, err = interactor.AccountRepository.FindById(identifier)
				if err != nil {
					return nil, err
				}
				return account, nil
			}
		}
		return nil, err
	}
	return nil, err
}
