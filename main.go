package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"orders-api/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=true"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("[!] Failed to connect to the database")
	}

	db.Exec("CREATE DATABASE orders_db")
	db.Exec("USE orders_db")

	db.AutoMigrate(&models.Order{}, &models.Item{})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")

	initDB()
	fmt.Println("[*] Starting server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	db.Create(&order)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["orderId"]
	id64, _ := strconv.ParseUint(id, 10, 64)
	orderId := uint(id64)

	var order models.Order
	db.Preload("Items").First(&order, orderId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	w.Header().Set("Content-Type", "application/json")
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {

}

func deleteOrder(w http.ResponseWriter, r *http.Request) {

}
