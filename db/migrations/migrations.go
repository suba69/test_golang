// migrations.go

package migrations

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ApplyMigrations(dbURL string) error {
	m, err := migrate.New("file://C:/test_golang/db/migrations", dbURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func CheckMigrationsStatus(dbURL string) error {
	m, err := migrate.New("file://C:/test_golang/db/migrations", dbURL)
	if err != nil {
		return err
	}

	version, dirty, err := m.Version()
	if err != nil {
		return err
	}

	if dirty {
		return errors.New("Migrations are seen as dirty")
	}

	log.Printf("Migration version: %d\n", version)

	return nil
}
