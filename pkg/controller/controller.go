package controllers
import(
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/DarkTHor55/go-bookstore/pkg/utils"
	"github.com/DarkTHor55/go-bookstore/pkg/models"
)
var NewBook models.Book 

func GetBook(w http.ResonseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResonseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId := vars["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error parsing",err)
	}
	bookDetails,_:=models.GetBookByBookId(bookId)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResonseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newBook :=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=newBook.CreateBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
func DeleteBook(w http.ResonseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars:=mux.Vars(r)
	bookId := vars["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
        fmt.Println("Error parsing",err)
    }
	book:=models.DelBook(Id)
	res,_:=json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResonseWriter, r *http.Request){
	newBook:=&models.Book{}
	w.Header().Set("Content-Type", "application/json")
    utils.ParseBody(r,newBook)
	v:= mux.Vars(r)
	bookId := v["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("Error parsing",err)
	}
	bookDetails,db:=models.GetBookById(Id)
	if newBook.name!=""{
		bookDetails.Name=newBook.Name

	} 
	if newBook.author!=""{
        bookDetails.Author=newBook.Author
    }
	if newBook.Publication!=""{
        bookDetails.Publication=newBook.Publication
    }
	db.save(&bookDetails)
	res,_:=json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}