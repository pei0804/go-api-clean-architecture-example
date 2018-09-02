package account

import "github.com/pei0804/go-api-clean-architecture-example/model"

type Repository interface {
	GetList() ([]model.Account, error)
	GetByID(id int) (*model.Account, error)
	Update(model.Account) (*model.Account, error)
	Add(*model.Account) (int, error)
	Delete(id int) (bool, error)
}
