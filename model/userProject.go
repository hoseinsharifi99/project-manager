package model

import (
	"github.com/jinzhu/gorm"
)

type UserProject struct {
	gorm.Model
	UserID    uint `gorm:"unique_index;not null"`
	ProjectID uint `gorm:"unique_index;not null"`
	Duration  uint
}
