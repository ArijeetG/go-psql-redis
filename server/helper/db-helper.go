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
	// psql := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	url := fmt.Sprintf("postgres://%v:%v@postgres:%v/%v?sslmode=disable",
		user, password, port, dbname)
	fmt.Println(url)
	Db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err.Error())
		return ConnectToDatabase()
	}

	log.Println("Successfully connected to db!!")
	return Db
	// defer db.Close()
}
