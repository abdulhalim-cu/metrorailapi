package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/abdulhalim-cu/metrorailapi/dbutils"
	_ "github.com/mattn/go-sqlite3"
)

// DB driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds the information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

func main() {

	db, err := sql.Open("sqlite3", "./rainapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(db)

}
