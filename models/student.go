package models

type Student struct {
	Id    string  `json:"id" binding:"required"` // Challenge 3: id must not be empty
	Name  string  `json:"name" binding:"required"` // Challenge 3: name must not be empty
	Major string  `json:"major"`
	GPA   float64 `json:"gpa" binding:"required,min=0,max=4"` // Challenge 3: gpa 0.00 - 4.00
}