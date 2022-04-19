package main

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func RunMigration() {
	db, err := sql.Open("postgres", "postgres://postgres:secret@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Panic(err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	m, er := migrate.NewWithDatabaseInstance(
		filepath.Dir("../migrations/initialise.sql"),
		"postgres", driver)
	if er != nil {
		log.Panic(er.Error())
	}
	m.Steps(2)
}
