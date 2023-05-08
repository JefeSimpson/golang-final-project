package repository

import (
	"final-project/model"
	"gorm.io/gorm"
)

type PurchaseCrud struct {
	db *gorm.DB
}

func NewPurchaseCrud(db *gorm.DB) *PurchaseCrud {
	return &PurchaseCrud{db: db}
}

func (repo *PurchaseCrud) Create(purchase model.Purchase) error {
	if err := repo.db.Create(&purchase).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PurchaseCrud) GetAll() ([]model.Purchase, error) {
	var purchases []model.Purchase

	if err := repo.db.Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}

func (repo *PurchaseCrud) GetById(id int) (model.Purchase, error) {
	var purchase model.Purchase

	if err := repo.db.First(&purchase, id).Error; err != nil {
		return model.Purchase{}, err
	}

	return purchase, nil
}

func (repo *PurchaseCrud) GetByItem(itemId int) ([]model.Purchase, error) {
	var purchases []model.Purchase

	if err := repo.db.Where("item_id = ?", itemId).Find(&purchases).Error; err != nil {
		return nil, err
	}
	return purchases, nil
}

func (repo *PurchaseCrud) GetByUser(userId int) ([]model.Purchase, error) {
	var purchases []model.Purchase

	if err := repo.db.Where("user_id = ?", userId).Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}

func (repo *PurchaseCrud) Update(purchase model.Purchase, id int) error {
	if err := repo.db.Model(&model.Purchase{}).Where("id = ?", id).Updates(purchase).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PurchaseCrud) Delete(id int) error {
	if err := repo.db.Delete(&model.Purchase{}, id).Error; err != nil {
		return err
	}
	return nil
}