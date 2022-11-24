package db

import (
	"database/sql"
	"log"

	"github.com/qustavo/dotsql"
)

func getQueryFile() *dotsql.DotSql {
	// Queries in separete file for ease of maintenance.
	dot, err := dotsql.LoadFromFile("pkg/db/queries.sql")
	if err != nil {
		log.Fatal("Cannot load queries.sql")
	}

	return dot
}

func getVersion(db *sql.DB) string {
	dot := getQueryFile()

	row, err := dot.QueryRow(db, "get-version")
	if err != nil {
		log.Fatal("Error when trying to connect to DB. Err:", err)
	}
	var version string
	row.Scan(&version)

	return version
}
