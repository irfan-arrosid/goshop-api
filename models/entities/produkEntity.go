package entities

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	Id             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama_produk    string         `gorm:"type:varchar(255);index" json:"nama_produk"`
	Slug           string         `gorm:"type:varchar(255)" json:"slug"`
	Harga_reseller string         `gorm:"type:varchar(255)" json:"harga_reseller"`
	Harga_konsumen string         `gorm:"type:varchar(255)" json:"harga_konsumen"`
	Stok           int            `json:"stok"`
	Deskripsi      string         `json:"deskripsi"`
	Id_toko        Toko           `gorm:"foreignKey:Id_toko" json:"Id_toko"`
	Id_category    Category       `gorm:"foreignKey:Id_category" json:"Id_category"`
	Created_at     time.Time      `json:"created_at"`
	Updated_at     time.Time      `json:"updated_at"`
	Deleted_at     gorm.DeletedAt `gorm:"index" json:"-"`
}
