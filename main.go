package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")

	fmt.Println("[*] Starting server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createOrder(w http.ResponseWriter, r *http.Request) {

}

func getOrder(w http.ResponseWriter, r *http.Request) {
	
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	
}