package model

import (
	"github.com/jinzhu/gorm"
)

type UserProject struct {
	gorm.Model
	UserId    uint
	ProjectId uint
	Duration  uint
}
