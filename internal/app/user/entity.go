package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"type:varchar(36);primaryKey"`
	Name      string
	Email     string
	Password  string
	Address   string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Id == "" {
		user.Id = uuid.New().String() // should be string in mysql
	}
	return nil
}
