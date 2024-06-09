package db

import (
	"database/sql"

	"github.com/wonyus/update-status-service/contexts"
)

type DBClient struct {
	*sql.DB
}

func NewPGDB(ctx *contexts.Resource) *DBClient {
	// Connect to database
	DB, err := sql.Open("postgres", ctx.DB_PG_URL)
	if err != nil {
		panic("Failed to connect to  db")
	}

	return &DBClient{DB}
}
