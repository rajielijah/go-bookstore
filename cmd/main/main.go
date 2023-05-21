package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rajielijah/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
