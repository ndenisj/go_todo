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

/***
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ndenisj/go_todo/utils"
)

var testQueries *Queries

// TestMain: main entrypoint of the test
func TestMain(m *testing.M) {
	//load config
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to DB")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

*/
