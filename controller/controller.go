package controller

import (
	"final-project/service"
	"github.com/gorilla/mux"
)

type Controller struct {
	services *service.Service
}

func NewController(services *service.Service) *Controller {
	return &Controller{services: services}
}

func (c *Controller) InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/login", c.login).Methods("POST")

	auth.HandleFunc("/register", c.register).Methods("POST")

	items := router.PathPrefix("/items").Subrouter()
	comments := router.PathPrefix("/comments").Subrouter()
	category := router.PathPrefix("/categories").Subrouter()
	users := router.PathPrefix("/users").Subrouter()
	purchase := router.PathPrefix("/purchases").Subrouter()

	// -> middleware

	items.Use(c.verification)
	comments.Use(c.verification)
	category.Use(c.verification)
	users.Use(c.verification)
	purchase.Use(c.verification)

	// -> category start

	category.HandleFunc("", c.getCategories).Methods("GET")

	category.HandleFunc("/{id}", c.getCategoryById).Methods("GET")

	category.HandleFunc("", c.addCategory).Methods("POST")

	category.HandleFunc("/{id}", c.updateCategory).Methods("PUT")

	category.HandleFunc("/{id}", c.deleteCategory).Methods("DELETE")

	// <- category end

	// -> comment start

	comments.HandleFunc("", c.getComments).Methods("GET")

	comments.HandleFunc("/{id}", c.getCommentById).Methods("GET")

	comments.HandleFunc("/items/{id}", c.getCommentByItemId).Methods("GET")

	comments.HandleFunc("/users/{id}", c.getCommentByUserId).Methods("GET")

	comments.HandleFunc("", c.addComment).Methods("POST")

	comments.HandleFunc("/{id}", c.updateComment).Methods("PUT")

	comments.HandleFunc("/{id}", c.deleteComment).Methods("DELETE")

	// <- comment end

	// -> purchase start

	purchase.HandleFunc("", c.getPurchases).Methods("GET")

	purchase.HandleFunc("/{id}", c.getPurchaseById).Methods("GET")

	purchase.HandleFunc("/items/{id}", c.getPurchasesByItemId).Methods("GET")

	purchase.HandleFunc("/users/{id}", c.getPurchasesByUserId).Methods("GET")

	purchase.HandleFunc("", c.addPurchase).Methods("POST")

	purchase.HandleFunc("/{id}", c.updatePurchase).Methods("PUT")

	purchase.HandleFunc("/{id}", c.deletePurchase).Methods("DELETE")

	// <- purchase end

	// -> item start

	items.HandleFunc("", c.getItems).Methods("GET")

	items.HandleFunc("/{id}", c.getItemById).Methods("GET")

	items.HandleFunc("/categories/{id}", c.getItemsByCategory).Methods("GET")

	items.HandleFunc("/sort/name", c.sortByName).Methods("GET")

	items.HandleFunc("/sort/rating", c.sortByRating).Methods("GET")

	items.HandleFunc("/sort/price", c.sortByPrice).Methods("GET")

	items.HandleFunc("", c.createItem).Methods("POST")

	items.HandleFunc("/{id}", c.updateItem).Methods("PUT")

	items.HandleFunc("/rate/{id}", c.setItemRating).Methods("PUT")

	items.HandleFunc("/{id}", c.deleteItemById).Methods("DELETE")

	// <- item end

	// -> user start

	users.HandleFunc("/balance/{id}", c.getBalance).Methods("GET")

	users.HandleFunc("/deposit/{id}", c.deposit).Methods("PUT")

	// <- user end

	return router
}
