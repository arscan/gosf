package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type postHandler struct{}

func (ph *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi There."))
}

// Gets all orders
func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Got an order"))
}

// Creates a new order
func PostOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted an order"))
}

// Updates an order
func PutOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Put an order"))

}

// Deletes an order
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleted an order"))
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/test", GetOrder).Methods("GET")
	mux.HandleFunc("/test", PostOrder).Methods("POST")
	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)

}
