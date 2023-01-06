package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/todo?sslmode=disable"
)

var testQueries *Queries

// TestMain: main entrypoint of the test
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to DB")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
