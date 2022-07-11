package model

import (
	"database/sql"
	"log"
	"os"
)

var con *sql.DB

func Connect() *sql.DB {
	// "root:1234@tcp:(localhost:6379)/mysql"
	dbName := os.Getenv("DB_NAME")
	dbString := os.Getenv("DB_STRING")
	db, err := sql.Open(dbName, dbString)
	if err != nil {
		log.Fatal(err)
	}
	println("MySQL Connected sucessfully...")
	con = db
	return db
}
