package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Connect to DB
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s schema=%s sslmode=disable",
		"postgres", "password", "aditum", "aps")
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
}
