package usecase

import (
	"github.com/pei0804/go-api-clean-architecture-example/account"
	"github.com/pei0804/go-api-clean-architecture-example/model"
)

type accountUsecase struct {
	accountRepo account.Repository
}

func NewAccountUsecase(accountRepo account.Repository) account.Usecase {
	return &accountUsecase{
		accountRepo: accountRepo,
	}
}

func (au *accountUsecase) GetList() ([]*model.Account, error) {
	return nil, nil
}
