package db

import (
	"database/sql"
	"fast_food_auth/config"
	"fmt"
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
		credentials := config.GetDatabaseCredentials();
		

		dbInstance, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", credentials.Username, credentials.Password, credentials.Host, credentials.Database))
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
