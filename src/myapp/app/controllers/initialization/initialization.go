package initialization

import "database/sql"

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	var err error
	if DB == nil {
		DB, err = sql.Open("postgres", connStr)
	}

	return DB, err
}
