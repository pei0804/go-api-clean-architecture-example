package account

import "github.com/pei0804/go-api-clean-architecture-example/model"

type Usecase interface {
	GetList() ([]*model.Account, error)
}
