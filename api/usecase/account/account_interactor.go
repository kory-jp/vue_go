package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/kory-jp/vue_go/api/domain"
)

type AccountInteractor struct {
	AccountRepository AccountRepository
}

func (interactor *AccountInteractor) Add(ac domain.Account) (account *domain.Account, err error) {
	if err = ac.AccountValidate(); err == nil {
		ac.Password = ac.Encrypt(ac.Password)
		identifier, err := interactor.AccountRepository.Store(ac)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("データ保存に失敗しました")
			return nil, err
		} else {
			account, err = interactor.AccountRepository.FindById(identifier)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				err = errors.New("データ取得に失敗しました")
				return nil, err
			}
			return account, nil
		}
	}
	return nil, err
}
