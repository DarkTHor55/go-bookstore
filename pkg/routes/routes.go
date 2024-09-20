package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/DarkTHor55/go-bookstore/pkg/controller"
)

func RegisterBookStoreRoutes() *mux.Router {
	router := mux.NewRouter()

	// Create a new book
	router.HandleFunc("/books", controller.CreateBook).Methods(http.MethodPost)

	// Get a single book by ID
	router.HandleFunc("/books/{id}", controller.GetBook).Methods(http.MethodGet)

	// Get all books
	router.HandleFunc("/books", controller.GetAllBooks).Methods(http.MethodGet)

	// Update a book by ID
	router.HandleFunc("/books/{id}", controller.UpdateBook).Methods(http.MethodPut)

	// Delete a book by ID
	router.HandleFunc("/books/{id}", controller.DeleteBook).Methods(http.MethodDelete)

	return router
}
