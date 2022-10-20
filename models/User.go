package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"not null; unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
