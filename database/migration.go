package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunManualMigration(dbURL string) {
	m, err := migrate.New(
		"file://database/migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	// m.Force(9)

	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	log.Println("completed db migration")
}
