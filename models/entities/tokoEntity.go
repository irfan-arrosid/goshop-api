package entities

import (
	"time"

	"gorm.io/gorm"
)

type Toko struct {
	Id         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user    uint           `gorm:"type:varchar(255);index" json:"id_user"`
	User       User           `gorm:"foreignKey:Id_user" json:"user"`
	Nama_toko  string         `gorm:"type:varchar(255)" json:"nama_toko"`
	Url_foto   string         `gorm:"type:varchar(255)" json:"url_foto"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `gorm:"index" json:"-"`
}
