package repository

import (
	"golang-tutorial/contract"	
	"golang-tutorial/entity"

	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func implAccountRepository(db *gorm.DB) contract.AccountRepository {
	return &AccountRepo{
		db: db,
	}
}

func (r *AccountRepo) GetAccount(id int) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Table("account").Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func (r *AccountRepo) CreateAccount(account *entity.Account) error {
	return r.db.Table("account").Create(account).Error
}

func (r *AccountRepo) GetAccountByEmail(email string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Table("account").Where("email = ?", email).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) CheckEmail(email string) (bool, error) {
	var exist bool
	err := r.db.Raw("SELECT EXISTS (SELECT 1 FROM account WHERE email = $1)", email).Scan(&exist).Error
	return exist, err
}
