package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/qustavo/dotsql"
)

func getQueryFile() *dotsql.DotSql {

	// When running tests, path changes, so check current envirome
	var path string
	if os.Getenv("ENVIROMENT") == "dev" {
		path = "pkg/db/queries.sql"
	} else {
		path = "../db/queries.sql"
	}

	// Queries in separete file for ease of maintenance.
	dot, err := dotsql.LoadFromFile(path)
	if err != nil {
		log.Fatal("Cannot load queries.sql")
	}

	return dot
}

func GetVersion(db *sql.DB) string {
	dot := getQueryFile()

	row, err := dot.QueryRow(db, "get-version")
	if err != nil {
		log.Fatal("Error when trying to connect to DB. Err:", err)
	}
	var version string
	row.Scan(&version)

	return version
}
