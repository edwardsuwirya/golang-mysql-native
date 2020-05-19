package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/enigma")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
