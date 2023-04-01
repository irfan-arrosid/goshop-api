package entities

import (
	"time"

	"gorm.io/gorm"
)

type (
	Produk struct {
		Id             uint         `gorm:"primaryKey;autoIncrement" json:"id"`
		Nama_produk    string       `gorm:"type:varchar(255);index" json:"nama_produk"`
		Slug           string       `gorm:"type:varchar(255)" json:"slug"`
		Harga_reseller string       `gorm:"type:varchar(255)" json:"harga_reseller"`
		Harga_konsumen string       `gorm:"type:varchar(255)" json:"harga_konsumen"`
		Stok           int          `json:"stok"`
		Deskripsi      string       `json:"deskripsi"`
		Id_toko        uint         `gorm:"type:varchar(255);index" json:"id_toko"`
		Toko           Toko         `gorm:"foreignKey:Id_toko" json:"toko"`
		Id_category    uint         `gorm:"type:varchar(255);index" json:"id_category"`
		Category       Category     `gorm:"foreignKey:Id_category" json:"category"`
		Foto           []FotoProduk `gorm:"foreignKey:Id_produk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	}

	FotoProduk struct {
		Id         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
		Id_produk  uint           `gorm:"type:varchar(255);index" json:"id_produk"`
		Produk     Produk         `gorm:"foreignKey:Id_produk" json:"produk"`
		Url        string         `gorm:"type:varchar(255)" json:"url"`
		Created_at time.Time      `json:"created_at"`
		Updated_at time.Time      `json:"updated_at"`
		Deleted_at gorm.DeletedAt `gorm:"index" json:"-"`
	}
)
