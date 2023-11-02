package repository

import (
	"gobank/config"
	"gobank/model"
	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	DB *gorm.DB
}

func NewAccountRepository() AccountRepository {
	DB, _ := config.NewMySQLDB()
	return &AccountRepositoryImpl{
		DB: DB,
	}
}

func (repository AccountRepositoryImpl) CreateAccount(account *model.Account) error {
	tx := repository.DB.Create(&account)
	return tx.Error

}

func (repository AccountRepositoryImpl) DeleteAccount(id int) error {
	var account model.Account
	return repository.DB.Delete(&account, id).Error
}

func (repository AccountRepositoryImpl) UpdateAccount(id int, account *model.Account) error {
	var currentAccount model.Account
	repository.DB.First(&currentAccount, id)

	if account.FirstName != "" {
		currentAccount.FirstName = account.FirstName
	}

	if account.LastName != "" {
		currentAccount.LastName = account.LastName
	}

	if account.Balance != 0 {
		currentAccount.Balance = account.Balance
	}

	tx := repository.DB.Save(&currentAccount)
	return tx.Error
}

func (repository AccountRepositoryImpl) GetAccountById(id int) (*model.Account, error) {
	var account model.Account
	repository.DB.First(&account, id)
	return &account, nil

}

func (repository AccountRepositoryImpl) GetAccounts() ([]*model.Account, error) {
	var accounts []*model.Account
	repository.DB.Find(&accounts)
	return accounts, nil
}
