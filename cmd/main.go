package main

import (
	"log"
	"net/http"
	"github.com/DarkTHor55/go-bookstore/pkg/router"
)

func main() {
	r := router.RegisterBookStoreRoutes()
	log.Fatal(http.ListenAndServe("localhost:8001", r))
}
