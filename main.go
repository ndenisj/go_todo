package main

import (
	"database/sql"
	"log"

	"github.com/ndenisj/go_todo/api"
	db "github.com/ndenisj/go_todo/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/todo?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
