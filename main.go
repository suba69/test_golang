package main

import (
	_ "github.com/mailgun/mailgun-go/v4"
	"log"
	"net/http"
	"test_golang/db"
	"test_golang/handler"
)

func main() {
	database := db.NewDB()
	defer database.Close()

	http.HandleFunc("/exchange-rate", handler.GetCurrentExchangeRate(database))
	http.HandleFunc("/subscribe", handler.SubscribeHandler(database))

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
