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
		dbInstance, err := sql.Open("postgres", "postgresql://user:pass@localhost/dbname?sslmode=disable")
		if err != nil {
			return
		}
		if err = dbInstance.Ping(); err != nil {
			dbInstance = nil
		}
		queries = New(dbInstance)
	})
	return queries, err
}
