package main

import (
	"database/sql"
	"log"

	"github.com/ndenisj/go_todo/api"
	db "github.com/ndenisj/go_todo/db/sqlc"
	"github.com/ndenisj/go_todo/utils"

	_ "github.com/lib/pq"
)

func main() {
	//load config
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to DB", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can not create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
