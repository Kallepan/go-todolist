package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// ErrNoMatch is returned when a query returns no rows in the result set.
var ErrNoMatch = fmt.Errorf("no matching record found")

type Database struct {
	CON *sql.DB
}

func InitDB(username, password, database, host string, port int) (Database, error) {
	db := Database{}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	dbCon, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}

	db.CON = dbCon
	err = db.CON.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Successfully connected to database! :D")

	return db, nil
}
