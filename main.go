package main

import (
	_ "github.com/mailgun/mailgun-go/v4"
	"log"
	"net/http"
	"test_golang/db"
	"test_golang/db/migrations"
	"test_golang/handler"
)

func main() {
	database := db.NewDB()
	defer database.Close()

	dbURL := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	if err := migrations.ApplyMigrations(dbURL); err != nil {
		log.Fatal(err)
	}

	if err := migrations.CheckMigrationsStatus(dbURL); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/exchange-rate", handler.GetCurrentExchangeRate(database))
	http.HandleFunc("/subscribe", handler.SubscribeHandler(database))

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
