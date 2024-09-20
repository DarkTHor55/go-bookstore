package router

import (
	"github.com/gorilla/mux"
	"github.com/DarkTHor55/go-bookstore/pkg/controller"
)

func RegisterBookStoreRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")

	return router
}
