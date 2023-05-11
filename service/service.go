package service

import (
	"final-project/model"
	"final-project/repository"
)

type Categories interface {
	Create(category model.Category) error
	GetAll() ([]model.Category, error)
	GetById(id int) (model.Category, error)
	Delete(id int) error
	Update(category model.Category, id int) error
}
type Comments interface {
	CreateComment(comment model.Comment) error
	GetComments() ([]model.Comment, error)
	GetCommentById(id int) (model.Comment, error)
	GetCommentsByItemId(itemId int) ([]model.Comment, error)
	GetCommentsByUserId(userId int) ([]model.Comment, error)
	DeleteComment(id int) error
	UpdateComment(comment model.Comment, id int) error
}
type Purchases interface {
	Create(purchase model.Purchase) error
	GetAll() ([]model.Purchase, error)
	GetById(id int) (model.Purchase, error)
	GetByItemId(itemId int) ([]model.Purchase, error)
	GetByUserId(userId int) ([]model.Purchase, error)
	Delete(id int) error
	Update(purchase model.Purchase, id int) error
}
type Items interface {
	Create(item model.Item) error
	GetAll() ([]model.Item, error)
	GetById(id int) (model.Item, error)
	GetByCategoryId(categoryId int) ([]model.Item, error)
	SortByName(name string) ([]model.Item, error)
	SortByRating(rating string) ([]model.Item, error)
	SortByPrice(price string) ([]model.Item, error)
	SetRating(rating float32, id int) error
	Delete(id int) error
	Update(item model.Item, id int) error
}
type Authorization interface {
	RegisterUser(user model.User) error
	GenerateToken(username, password string) (string, string, error)
	ParseToken(accessToken string) (string, error)
	RefreshToken(accessToken string) (string, string, error)
}

type Users interface {
	Deposit(id int, dep float32) error
	GetBalance(id int) (float32, error)
}

type Service struct {
	Categories
	Comments
	Purchases
	Items
	Authorization
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Categories:    NewCategoryService(repo.CategoryRepo),
		Comments:      NewCommentService(repo.CommentRepo),
		Purchases:     NewPurchaseService(repo.PurchaseRepo, repo.UserRepo, repo.ItemRepo),
		Items:         NewItemService(repo.ItemRepo),
		Authorization: NewAuthService(repo.AuthorizationRepo),
		Users:         NewUserService(repo.UserRepo),
	}
}
