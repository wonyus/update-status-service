package initials

import (
	"database/sql"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := os.Getenv("DB_URL")
	// Connect to database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to  db")
	}
}
