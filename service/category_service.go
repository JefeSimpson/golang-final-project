package service

import (
	"final-project/model"
	"final-project/repository"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(repo repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category model.Category) error {
	if err := s.repo.Create(category); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (s *CategoryService) GetById(id int) (model.Category, error) {
	category, err := s.repo.GetById(id)
	if err != nil {
		return model.Category{}, err
	}
	return category, err
}

func (s *CategoryService) Update(category model.Category, id int) error {
	return s.repo.Update(category, id)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
