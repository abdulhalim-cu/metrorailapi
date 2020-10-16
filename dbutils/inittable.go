package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	stmt, err := dbDriver.Prepare(train)
	if err != nil {
		log.Println(err)
	}
	// create train table
	_, err = stmt.Exec()
	if err != nil {
		log.Println("Table already exists!")
	}
	stmt, _ = dbDriver.Prepare(station)
	stmt.Exec()
	stmt, _ = dbDriver.Prepare(schedule)
	stmt.Exec()
	log.Println("All tables created/initialized successfully")
}
