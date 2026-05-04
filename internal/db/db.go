package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./caselog.db")
	if err != nil {
		return nil, err
	}

	if err := createProjectsTable(database); err != nil {
		return nil, err
	}

	return database, nil
}

func createProjectsTable(database *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		company TEXT,
		main_skill TEXT,
		unit_price INTEGER,
		remote_type TEXT,
		team_size TEXT,
		status TEXT NOT NULL DEFAULT 'considering',
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := database.Exec(query)
	return err
}