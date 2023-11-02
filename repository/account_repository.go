package repository

import (
	"gobank/model"
)

type AccountRepository interface {
	CreateAccount(*model.Account) error
	DeleteAccount(int) error
	UpdateAccount(int, *model.Account) error
	GetAccountById(int) (*model.Account, error)
	GetAccounts() ([]*model.Account, error)
}
