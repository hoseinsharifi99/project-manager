package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	ProjectID   uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Duration    uint
	Description string
	UserProject []UserProject `gorm:"foreignkey:task_id"`
}
