package model

import (
	"database/sql"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp:(localhost:6379)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	println("MySQL Connected sucessfully...")
	con = db
	return db
}
