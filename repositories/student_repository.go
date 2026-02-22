package repositories

import (
	"database/sql"
	"errors"
	"go-api-gin/models"
)

// สร้าง Struct ที่เก็บการเชื่อมต่อ Database
type StudentRepository struct {
	db *sql.DB
}

// ฟังก์ชันสำหรับสร้าง Repository
func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

// GET All
func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.db.Query("SELECT id, name, major, gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
		students = append(students, s)
	}
	return students, nil
}

// GET by ID
func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	var s models.Student
	err := r.db.QueryRow("SELECT id, name, major, gpa FROM students WHERE id = ?", id).Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
	if err == sql.ErrNoRows {
		return nil, errors.New("student not found")
	}
	return &s, err
}

// POST
func (r *StudentRepository) Create(s models.Student) error {
	_, err := r.db.Exec("INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)", s.Id, s.Name, s.Major, s.GPA)
	return err
}

// PUT (Challenge 1)
func (r *StudentRepository) Update(id string, s models.Student) error {
	result, err := r.db.Exec("UPDATE students SET name=?, major=?, gpa=? WHERE id=?", s.Name, s.Major, s.GPA, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("student not found")
	}
	return nil
}

// DELETE (Challenge 2)
func (r *StudentRepository) Delete(id string) error {
	result, err := r.db.Exec("DELETE FROM students WHERE id=?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("student not found")
	}
	return nil
}