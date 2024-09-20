package models
import (
	"github.com/dARKthOR55/go-bookstore/pkg/config"
	"gorm.io/gorm"
)
var db *gorm.db
type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"
	Publication string `json:"publication"`
}
func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}
func(b *Book)CreateBook()*Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks()[]Book{
	var books []Book
	db.Find(&Books)
	return books
}
func GetBookById(Id int64 )*Book,gorm.DB{
	var getBook Book
	db:=db.where("id=?",Id).Find(&getBook)
	return &getBook,db
}
func delBook(Id int64)Book{
	var book Book
	db.where("id=?",Id).Delete(&book)
	return book
}
