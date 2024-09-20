package models

import (
	"github.com/DarkTHor55/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB 

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Publication string `gorm:"type:varchar(100)" json:"publication"`
}

func init() {
	config.Connect() 
	db = config.GetDB() 
	db.AutoMigrate(&Book{}) 
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books) 
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) { 
	var getBook Book
	db := db.Where("id = ?", Id).Find(&getBook) 
	return &getBook, db
}

func DeleteBook(Id int64) Book { 
	var book Book
	db.Where("id = ?", Id).Delete(&book) 
	return book
}
