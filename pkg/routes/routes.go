package router

import (
	"github.com/gorilla/mux"
	"github.com/DarkTHor55/go-bookstore/pkg/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
