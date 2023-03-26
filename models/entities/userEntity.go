package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama          string         `gorm:"type:varchar(255)" json:"nama"`
	Kata_sandi    string         `gorm:"column:password" json:"kata_sandi"`
	Notelp        string         `gorm:"type:varchar(255);uniqueIndex" json:"notelp"`
	Tanggal_lahir time.Time      `gorm:"type:date" json:"tanggal_lahir"`
	Jenis_kelamin string         `gorm:"type:varchar(255)" json:"jenis_kelamin"`
	Tentang       string         `json:"tentang"`
	Pekerjaan     string         `gorm:"type:varchar(255)" json:"pekerjaan"`
	Email         string         `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Id_provinsi   string         `gorm:"type:varchar(255)" json:"id_provinsi"`
	Id_kota       string         `gorm:"type:varchar(255)" json:"id_kota"`
	IsAdmin       bool           `json:"isAdmin"`
	Created_at    time.Time      `json:"created_at"`
	Updated_at    time.Time      `json:"updated_at"`
	Deleted_at    gorm.DeletedAt `gorm:"index" json:"-"`
}
