package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book used as a main REST resource
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {

	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr.Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr.Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr.Router", Year: "2012"},
		Book{ID: 4, Title: "Golang Concurrency", Author: "Mr.Concurrency", Year: "2013"},
		Book{ID: 5, Title: "Golang Good Parts", Author: "Mr.Good", Year: "2014"},
	)

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", RemoveBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all books")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets one book")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adds one book")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates one book")
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removes one book")
}
