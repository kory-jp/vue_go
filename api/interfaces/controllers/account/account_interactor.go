package controllers

import domain "github.com/kory-jp/vue_go/api/domain/account"

type AccountInteractor interface {
	Add(domain.Account) (*domain.Account, error)
}
