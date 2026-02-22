package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite", "students.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// สร้างตารางเตรียมไว้
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id TEXT PRIMARY KEY,
		name TEXT,
		major TEXT,
		gpa REAL
	)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	return db
}