package db

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq" // Postgres driver
)

var (
	queries *Queries
	dbOnce     sync.Once
)

func GetDBInstance() (*Queries, error) {
	var err error
	dbOnce.Do(func() {
		dbInstance, err := sql.Open("postgres", "postgresql://myuser:mypassword@localhost/mydatabase?sslmode=disable")
		if err != nil {
			panic(err)
		}
		if err = dbInstance.Ping(); err != nil {
			panic(err)
		}
		queries = New(dbInstance)
	})
	return queries, err
}
