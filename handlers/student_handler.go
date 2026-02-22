package handlers

import (
	"go-api-gin/models"
	"go-api-gin/services" // เปลี่ยนมาใช้ package services
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *services.StudentService // เปลี่ยนจาก repo เป็น service
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// (โค้ดฟังก์ชัน GET, POST, PUT, DELETE ข้างในยังเหมือนเดิม แค่เปลี่ยน h.repo เป็น h.service)

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.service.GetAllStudents() // <--- แก้ตรงนี้
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.service.GetStudentByID(id) // <--- แก้ตรงนี้
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var s models.Student
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: id, name required and GPA 0-4"})
		return
	}

	if err := h.service.CreateStudent(s); err != nil { // <--- แก้ตรงนี้
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, s)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var s models.Student
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed"})
		return
	}

	if err := h.service.UpdateStudent(id, s); err != nil { // <--- แก้ตรงนี้
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	s.Id = id
	c.JSON(http.StatusOK, s)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteStudent(id); err != nil { // <--- แก้ตรงนี้
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.Status(http.StatusNoContent)
}