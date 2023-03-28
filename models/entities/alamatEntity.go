package entities

import (
	"time"

	"gorm.io/gorm"
)

type Alamat struct {
	Id            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user       uint           `gorm:"type:varchar(255);index" json:"id_user"`
	User          User           `gorm:"foreignKey:Id_user" json:"user"`
	Judul_alamat  string         `gorm:"type:varchar(255)" json:"judul_alamat"`
	Nama_penerima string         `gorm:"type:varchar(255)" json:"nama_penerima"`
	Notelp        string         `gorm:"type:varchar(255)" json:"notelp"`
	Detail_alamat string         `gorm:"type:varchar(255)" json:"detail_alamat"`
	Created_at    time.Time      `json:"created_at"`
	Updated_at    time.Time      `json:"updated_at"`
	Deleted_at    gorm.DeletedAt `gorm:"index" json:"-"`
}
