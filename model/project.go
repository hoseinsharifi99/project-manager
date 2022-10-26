package model

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name        string        `gorm:"unique_index;not null"`
	UserProject []UserProject `gorm:"foreignkey:project_id"`
	Task        []Task        `gorm:"foreignkey:project_id"`
}
