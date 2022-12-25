package main

import (
	"fmt"
	"log"
	"net/http"
	"orders-api/models"

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

}

func getOrder(w http.ResponseWriter, r *http.Request) {

}

func getOrders(w http.ResponseWriter, r *http.Request) {

}

func updateOrder(w http.ResponseWriter, r *http.Request) {

}

func deleteOrder(w http.ResponseWriter, r *http.Request) {

}
