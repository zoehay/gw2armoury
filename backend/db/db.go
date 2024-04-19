package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func Init(connStr string) {
	var err error
	db, err = sql.Open("postgres", connStr) 
	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database", err)
	}
}

func applyMigrations(db *sql.DB,migrationsPath string, migrationConnString string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	return nil
}