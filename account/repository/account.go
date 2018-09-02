package repository

import (
	"database/sql"

	"github.com/pei0804/go-api-clean-architecture-example/account"
	"github.com/pei0804/go-api-clean-architecture-example/model"
)

type accountRepository struct {
	Conn *sql.DB
}

func NewAccountRepository(Conn *sql.DB) account.Repository {
	return &accountRepository{Conn: Conn}
}

func (ar *accountRepository) GetList() ([]model.Account, error) {
	return nil, nil
}

func (ar *accountRepository) GetByID(id int) (*model.Account, error) {
	return nil, nil
}

func (ar *accountRepository) Update(a model.Account) (*model.Account, error) {
	return nil, nil
}

func (ar *accountRepository) Add(a *model.Account) (int, error) {
	return 1, nil
}

func (ar *accountRepository) Delete(id int) (bool, error) {
	return false, nil
}
