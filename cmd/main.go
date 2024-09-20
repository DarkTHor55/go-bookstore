package main

import (
    "log"
    "net/http"
	"github.com/DarkTHor55/go-bookstore/pkg/routers"
	"gorm.io/driver/mysql"
	"github.com/gorilla/mux"
)
func main() {
	r:=mux.NewRouter()
	routers.RegisterBookStoreRouter(r)
	log.fatal(http.ListenAndServe("localhost:8000",r))
}