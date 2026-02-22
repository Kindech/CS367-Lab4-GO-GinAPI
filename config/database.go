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

	// 1. สร้างตารางเตรียมไว้
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

	db.Exec(`
	INSERT OR IGNORE INTO students (id, name, major, gpa) VALUES
	('66090001', 'John', 'Computer Science', 3.80),
	('66090002', 'Jane', 'Information Technology', 3.50),
	('66090003', 'Mike', 'Software Engineering', 3.65)
	`)

	return db
}
