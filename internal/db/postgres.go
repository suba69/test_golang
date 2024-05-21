package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type Database struct {
	*sql.DB
}

type DatabaseInterface interface {
	InsertExchangeRate(rate float64) error
}

func NewDB() *Database {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{db}
}

func (database *Database) Close() {
	if database != nil && database.DB != nil {
		database.DB.Close()
	}
}

func (database *Database) InsertExchangeRate(rate float64) error {
	_, err := database.Exec("INSERT INTO exchange_rates (rate) VALUES ($1)", rate)
	return err
}

func (database *Database) CheckEmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM subscribed_users WHERE email = $1)`
	var exists bool
	err := database.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (database *Database) InsertSubscribedUser(email string) error {
	query := `INSERT INTO subscribed_users (email) VALUES ($1)`
	_, err := database.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLastEmailSentTime оновлює час останнього відправленого листа для підписаного користувача
func (database *Database) UpdateLastEmailSentTime(email string) error {
	query := `UPDATE subscribed_users SET last_email_sent_at = $1 WHERE email = $2`
	_, err := database.Exec(query, time.Now(), email)
	if err != nil {
		return err
	}
	return nil
}

// GetSubscribedUsers отримує список підписаних користувачів з бази даних
func (db *Database) GetSubscribedUsers() ([]string, error) {
	// Підготовка запиту до бази даних
	query := `SELECT email FROM subscribed_users`

	// Виконання запиту та отримання результатів
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Створення змінної для зберігання списку електронних адрес
	var emails []string

	// Ітерація через результати запиту
	for rows.Next() {
		var email string
		// Сканування значень рядка та збереження їх у змінну
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		// Додавання електронної адреси до списку
		emails = append(emails, email)
	}
	// Перевірка помилок, які можуть виникнути під час ітерації по результатам запиту
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Повернення списку електронних адрес
	return emails, nil
}
