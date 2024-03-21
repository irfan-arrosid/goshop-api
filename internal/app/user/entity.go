package user

import (
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Address  string
	IsAdmin  bool
	gorm.Model
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	user.Id = uuid.New()
// 	return nil
// }
