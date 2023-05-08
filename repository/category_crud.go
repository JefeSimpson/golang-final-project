package repository

import (
	"final-project/model"
	"gorm.io/gorm"
)

type CategoryCrud struct {
	db *gorm.DB
}

func NewCategoryCrud(db *gorm.DB) *CategoryCrud {
	return &CategoryCrud{db: db}
}

func (repo *CategoryCrud) Create(category model.Category) error {
	if err := repo.db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func (repo *CategoryCrud) GetAll() ([]model.Category, error) {
	var categories []model.Category

	if err := repo.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (repo *CategoryCrud) GetById(id int) (model.Category, error) {
	var category model.Category

	if err := repo.db.First(&category, id).Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (repo *CategoryCrud) Update(category model.Category, id int) error {
	var oldCategory model.Category

	if err := repo.db.First(&oldCategory, id).Error; err != nil {
		return err
	}

	category.ID = oldCategory.ID

	if err := repo.db.Save(&category).Error; err != nil {
		return err
	}

	return nil
}

func (repo *CategoryCrud) Delete(id int) error {
	if err := repo.db.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
