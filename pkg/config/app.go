package config

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	db *gorm.DB
)
func connect(){
	// d,err:=gorm.Open("mysql","root:Rohit@123/bookstore?charset=utf&&parseTime=true&loc=Local")
	dsn := "root:Rohit@123@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
        panic("Failed to connect to DB")
    }
}
func GetDB() *gorm.DB{ 
	return db
}