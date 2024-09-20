package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DarkTHor55/go-bookstore/pkg/models"
	"github.com/DarkTHor55/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

// CreateBook handler
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// GetBook handler
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book ID"))
		return
	}
	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetAllBooks handler
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook handler
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book ID"))
		return
	}
	deletedBook := models.DeleteBook(bookId)
	res, _ := json.Marshal(deletedBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook handler
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book ID"))
		return
	}
	bookDetails, db := models.GetBookById(bookId)
	if bookDetails != nil {
		utils.ParseBody(r, bookDetails)
		db.Save(bookDetails)
		res, _ := json.Marshal(bookDetails)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
	}
}
