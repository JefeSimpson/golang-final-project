package service

import "final-project/repository"

type UserService struct {
	user repository.UserRepo
}

func NewUserService(user repository.UserRepo) *UserService {
	return &UserService{user: user}
}

func (s *UserService) Deposit(id int, dep float32) error {
	if err := s.user.Deposit(id, dep); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetBalance(id int) (float32, error) {
	balance, err := s.user.GetBalance(id)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
