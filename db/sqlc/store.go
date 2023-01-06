package db

import "database/sql"

// Store: provides all functions to execute db queries (it will allow me interact with the db)
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore: creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
