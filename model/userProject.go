package model

import (
	"github.com/jinzhu/gorm"
)

type UserProject struct {
	gorm.Model
	UserID    uint
	ProjectID uint
	Duration  uint
}
