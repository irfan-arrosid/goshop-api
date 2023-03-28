package entities

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama_category string         `gorm:"type:varchar(255)" json:"nama_category"`
	Created_at    time.Time      `json:"created_at"`
	Updated_at    time.Time      `json:"updated_at"`
	Deleted_at    gorm.DeletedAt `gorm:"index" json:"-"`
}
