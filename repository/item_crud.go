package repository

import (
	"final-project/model"
	"gorm.io/gorm"
)

type ItemCrud struct {
	db *gorm.DB
}

func NewItemCrud(db *gorm.DB) *ItemCrud {
	return &ItemCrud{db: db}
}

func (repo *ItemCrud) Create(item model.Item) error {
	item.RatingTotal += item.Rating
	item.RatingCount++
	if err := repo.db.Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ItemCrud) GetAll() ([]model.Item, error) {
	var items []model.Item

	if err := repo.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *ItemCrud) GetById(id int) (model.Item, error) {
	var item model.Item

	if err := repo.db.First(&item, id).Error; err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (repo *ItemCrud) GetItemByCategoryId(categoryId int) ([]model.Item, error) {
	var items []model.Item

	if err := repo.db.Where("category_id = ?", categoryId).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *ItemCrud) SortByName(sort string) ([]model.Item, error) {
	var items []model.Item

	if err := repo.db.Where("UPPER(name) LIKE UPPER(?)", "%"+sort+"%").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *ItemCrud) SortByRating(sort string) ([]model.Item, error) {
	var items []model.Item

	if err := repo.db.Order("rating " + sort).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *ItemCrud) SortByPrice(sort string) ([]model.Item, error) {
	var items []model.Item

	if err := repo.db.Order("price " + sort).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *ItemCrud) Update(item model.Item, id int) error {
	var oldItem model.Item

	if err := repo.db.First(&oldItem, id).Error; err != nil {
		return err
	}

	item.ID = oldItem.ID

	if err := repo.db.Save(&item).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ItemCrud) SetRating(rating float32, id int) error {
	var item model.Item

	if err := repo.db.First(&item, id).Error; err != nil {
		return err
	}

	item.RatingTotal += rating
	item.RatingCount++

	item.Rating = item.RatingTotal / float32(item.RatingCount)

	if err := repo.db.Save(&item).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ItemCrud) Delete(id int) error {
	if err := repo.db.Delete(&model.Item{}, id).Error; err != nil {
		return err
	}

	return nil
}
