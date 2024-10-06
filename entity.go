package main

import "github.com/google/uuid"

//tabel database
type Student struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	GPA         float64   `json:"gpa"`
	IsGraduate  bool      `json:"is_graduate"`
}
