package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("error creating tables: %w", err)
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS debts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		total_amount REAL NOT NULL,
		remaining_amount REAL NOT NULL,
		due_date DATETIME NOT NULL,
		interest_rate REAL NOT NULL,
		num_installments INTEGER NOT NULL,
		installment_amount REAL NOT NULL,
		payment_day INTEGER NOT NULL,
		paid BOOLEAN DEFAULT 0,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS incomes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		source TEXT NOT NULL,
		date DATETIME NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS payments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		debt_id INTEGER NOT NULL,
		receipt_filename TEXT,
		date DATETIME NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		FOREIGN KEY (debt_id) REFERENCES debts(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	`

	_, err := db.Exec(schema)
	return err
}
