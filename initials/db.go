package initials

import (
	"database/sql"
	"os"

	"github.com/wonyus/update-status-service/utils"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := utils.Strip(os.Getenv("DB_URL"))
	// Connect to database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to  db")
	}
}
