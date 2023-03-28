package entities

import (
	"time"

	"gorm.io/gorm"
)

type FotoProduk struct {
	Id         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_produk  Produk         `gorm:"foreignKey:Id_produk" json:"Id_produk"`
	Url        string         `gorm:"type:varchar(255)" json:"url"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `gorm:"index" json:"-"`
}
