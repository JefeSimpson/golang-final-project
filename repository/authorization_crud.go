package repository

import (
	"errors"
	"final-project/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthorizationCrud struct {
	db *gorm.DB
}

func NewAuthorizationCrud(db *gorm.DB) *AuthorizationCrud {
	return &AuthorizationCrud{db: db}
}

func (repo *AuthorizationCrud) Register(user model.User) error {
	var count int64
	repo.db.Model(&model.User{}).Where("username = ?", user.Username).Count(&count)

	if count != 0 {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &model.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Balance:  0,
	}

	if err := repo.db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (repo *AuthorizationCrud) Login(username, password string) (model.User, error) {
	var user model.User

	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.User{}, err
	}

	return user, nil
}
