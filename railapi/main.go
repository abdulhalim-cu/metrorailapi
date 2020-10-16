package main

import (
	"database/sql"
	"log"

	"github.com/abdulhalim-cu/metrorailapi/dbutils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./rainapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(db)

}
