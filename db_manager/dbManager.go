package db_manager

import (
	"github.com/jinzhu/gorm"
)

type DbManager struct {
	db *gorm.DB
}

func NewDbManager(database *gorm.DB) *DbManager {
	return &DbManager{db: database}
}
