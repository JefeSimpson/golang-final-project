package service

import (
	"final-project/model"
	"final-project/repository"
)

type PurchaseService struct {
	repo repository.PurchaseRepo
	user repository.UserRepo
	item repository.ItemRepo
}

func NewPurchaseService(repo repository.PurchaseRepo,
	user repository.UserRepo,
	item repository.ItemRepo) *PurchaseService {
	return &PurchaseService{repo: repo, user: user, item: item}
}

func (s *PurchaseService) Create(purchase model.Purchase) error {

	item, err := s.item.GetById(purchase.ItemID)
	if err != nil {
		return err
	}

	err = s.user.Withdraw(purchase.UserID, item.Price)
	if err != nil {
		return err
	}

	if err := s.repo.Create(purchase); err != nil {
		err1 := s.user.Deposit(purchase.UserID, item.Price)
		if err1 != nil {
			return err1
		}
		return err
	}
	return nil
}

func (s *PurchaseService) GetAll() ([]model.Purchase, error) {
	comments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (s *PurchaseService) GetById(id int) (model.Purchase, error) {
	comment, err := s.repo.GetById(id)
	if err != nil {
		return model.Purchase{}, err
	}
	return comment, err
}

func (s *PurchaseService) GetByItemId(itemId int) ([]model.Purchase, error) {
	comments, err := s.repo.GetByItem(itemId)
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (s *PurchaseService) GetByUserId(userId int) ([]model.Purchase, error) {
	comments, err := s.repo.GetByUser(userId)
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (s *PurchaseService) Update(purchase model.Purchase, id int) error {
	// Update the item
	if err := s.repo.Update(purchase, id); err != nil {
		return err
	}
	return nil
}

func (s *PurchaseService) Delete(id int) error {
	return s.repo.Delete(id)
}
