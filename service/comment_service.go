package service

import (
	"final-project/model"
	"final-project/repository"
)

type CommentService struct {
	repo repository.CommentRepo
}

// NewCommentService возвращает новый экземпляр CommentService.
func NewCommentService(repo repository.CommentRepo) *CommentService {
	return &CommentService{repo}
}

// CreateComment создает новый комментарий.
func (s *CommentService) CreateComment(comment model.Comment) error {
	if err := s.repo.Create(comment); err != nil {
		return err
	}
	return nil
}

// GetComments возвращает все комментарии.
func (s *CommentService) GetComments() ([]model.Comment, error) {
	comments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return comments, err
}

// GetCommentById возвращает комментарий по его ID.
func (s *CommentService) GetCommentById(id int) (model.Comment, error) {
	comment, err := s.repo.GetById(id)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, err
}

// GetCommentsByItemId возвращает все комментарии для определенного товара.
func (s *CommentService) GetCommentsByItemId(itemId int) ([]model.Comment, error) {
	comments, err := s.repo.GetByItem(itemId)
	if err != nil {
		return nil, err
	}
	return comments, err
}

// GetCommentsByUserId возвращает все комментарии для определенного пользователя.
func (s *CommentService) GetCommentsByUserId(userId int) ([]model.Comment, error) {
	comments, err := s.repo.GetByUser(userId)
	if err != nil {
		return nil, err
	}
	return comments, err
}

// UpdateComment обновляет существующий комментарий.
func (s *CommentService) UpdateComment(comment model.Comment, id int) error {
	// Update the item
	if err := s.repo.Update(comment, id); err != nil {
		return err
	}
	return nil
}

// DeleteComment удаляет комментарий по его ID.
func (s *CommentService) DeleteComment(id int) error {
	return s.repo.Delete(id)
}
