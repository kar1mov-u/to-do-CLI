package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}
	createTableQuery := `
			CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed BOOLEAN,
			completed_at DATETIME
			);`
	if _, err = DB.Exec(createTableQuery); err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

}
