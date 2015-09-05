package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type postHandler struct{}

var orders = Orders{
	Order{Id: 1},
	Order{Id: 2},
	Order{Id: 3},
}

func (ph *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi There."))
}

// Gets all orders
func GetOrders(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		panic(err)
	}
}

// Get one
func GetOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	for _, o := range orders {
		if id, err := strconv.Atoi(vars["orderId"]); o.Id == id && err == nil {

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(o); err != nil {
				panic(err)
			}
			return

		}
	}

	w.Write([]byte("No such order"))

}

// Creates a new order
func PostOrder(w http.ResponseWriter, r *http.Request) {

	// find the order

	orders = append(orders, Order{Id: len(orders)})

	w.Write([]byte(fmt.Sprintf("Posted an order.  Now there are %d orders", len(orders))))
}

func PutOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, o := range orders {
		if id, err := strconv.Atoi(vars["orderId"]); o.Id == id && err == nil {

			w.Write([]byte(fmt.Sprintf("Put an order %d", id)))

			return

		}
	}

	w.Write([]byte("No such order"))

}

// Deletes an order
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for i, o := range orders {
		if id, err := strconv.Atoi(vars["orderId"]); o.Id == id && err == nil {

			orders = orders[:i+copy(orders[i:], orders[i+1:])]

			w.Write([]byte(fmt.Sprintf("Deleted an order %d", id)))

			return

		}
	}

	w.Write([]byte("No such order"))
}

func main() {
	mux := mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/orders", GetOrders).Methods("GET")
	mux.HandleFunc("/orders/{orderId}/", GetOrder).Methods("GET")
	mux.HandleFunc("/orders/", PostOrder).Methods("POST")
	mux.HandleFunc("/orders/{orderId}/", DeleteOrder).Methods("DELETE")
	mux.HandleFunc("/orders/{orderId}/", PutOrder).Methods("PUT")

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
