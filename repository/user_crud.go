package repository

import (
	"errors"
	"final-project/model"
	"gorm.io/gorm"
)

type UserCrud struct {
	db *gorm.DB
}

func NewUserCrud(db *gorm.DB) *UserCrud {
	return &UserCrud{db: db}
}

func (repo *UserCrud) GetBalance(id int) (float32, error) {
	var user model.User

	if err := repo.db.First(&user, id).Error; err != nil {
		return 0, err
	}

	return user.Balance, nil
}

func (repo *UserCrud) Withdraw(id int, price float32) error {
	var user model.User

	if err := repo.db.First(&user, id).Error; err != nil {
		return err
	}

	if user.Balance < price {
		return errors.New("not enough balance")
	}

	user.Balance -= price

	if err := repo.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserCrud) Deposit(id int, dep float32) error {
	var user model.User

	if err := repo.db.First(&user, id).Error; err != nil {
		return err
	}

	user.Balance += dep

	if err := repo.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
