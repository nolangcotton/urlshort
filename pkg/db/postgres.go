package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	port   = 5432
	user   = "admin"
	dbname = "urlshort"
	host   = "localhost"
)

func Conn() *sql.DB {
	password := os.Getenv("PG_PASS")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}
