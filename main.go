package main

import (
	"log"
	"net/http"
	"test_golang/handler"
	"test_golang/internal/db"
	"test_golang/internal/db/migrations"
)

func main() {
	database := db.NewDB()
	defer database.Close()

	dbURL := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	if err := migrations.ApplyMigrations(dbURL); err != nil {
		log.Fatal(err)
	}

	// Виправлено ім'я функції та аргументи
	if err := migrations.CheckMigrationsStatus(dbURL); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/exchange-rate", handler.GetCurrentExchangeRate(database, &handler.ExchangeRateAPIProvider{}))
	http.HandleFunc("/subscribe", handler.SubscribeHandler(database))

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
