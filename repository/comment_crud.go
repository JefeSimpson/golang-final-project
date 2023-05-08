package repository

import (
	"final-project/model"
	"gorm.io/gorm"
)

type CommentCrud struct {
	db *gorm.DB
}

func NewCommentCrud(db *gorm.DB) *CommentCrud {
	return &CommentCrud{db: db}
}

func (repo *CommentCrud) Create(comment model.Comment) error {
	if err := repo.db.Create(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (repo *CommentCrud) GetAll() ([]model.Comment, error) {
	var comments []model.Comment

	if err := repo.db.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (repo *CommentCrud) GetById(id int) (model.Comment, error) {
	var comment model.Comment

	if err := repo.db.First(&comment, id).Error; err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

func (repo *CommentCrud) GetByItem(itemId int) ([]model.Comment, error) {
	var comments []model.Comment

	if err := repo.db.Where("item_id = ?", itemId).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (repo *CommentCrud) GetByUser(userId int) ([]model.Comment, error) {
	var comments []model.Comment

	if err := repo.db.Where("user_id = ?", userId).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (repo *CommentCrud) Update(comment model.Comment, id int) error {
	var oldComment model.Comment

	if err := repo.db.First(&oldComment, id).Error; err != nil {
		return err
	}

	comment.ID = oldComment.ID

	if err := repo.db.Save(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (repo *CommentCrud) Delete(id int) error {
	if err := repo.db.Delete(&model.Comment{}, id).Error; err != nil {
		return err
	}

	return nil
}