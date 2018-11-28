package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//book model/struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"firstname"`
}

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//create new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

// main function
func main() {
	//init router
	router := mux.NewRouter()

	// route handler
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
