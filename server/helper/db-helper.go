package helper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	const (
		port     = 5432
		user     = "postgres"
		password = "secret"
		dbname   = "test"
	)
	url := fmt.Sprintf("postgres://%v:%v@host.docker.internal:%v/%v?sslmode=disable",
		user, password, port, dbname)
	fmt.Println(url)
	Db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err.Error())
		return ConnectToDatabase()
	}

	log.Println("Successfully connected to db!!")
	return Db
}
