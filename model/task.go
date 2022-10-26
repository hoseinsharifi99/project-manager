package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	ProjectID   uint   `gorm:"unique_index;not null"`
	Name        string `gorm:"not null"`
	Duration    uint
	Description string
}
