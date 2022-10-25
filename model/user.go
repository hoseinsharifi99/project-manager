package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique_index;not null"`
	Password    string `gorm:"not null"`
	Admin       int
	UserProject []UserProject `gorm:"foreignkey:user_id"`
}

func HashPassword(pass string) (string, error) {
	if len(pass) <= 0 {
		return "", errors.New("pass cant be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash), err
}

func (user *User) ValidatePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}
