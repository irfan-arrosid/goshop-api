package entities

import (
	"time"

	"gorm.io/gorm"
)

type DetailTrx struct {
	Id            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_trx        string         `gorm:"type:varchar(255);index" json:"id_trx"`
	Trx           Trx            `gorm:"foreignKey:Id_trx" json:"trx"`
	Id_log_produk uint           `gorm:"type:varchar(255);index" json:"Id_log_produk"`
	Log_produk    LogProduk      `gorm:"foreignKey:Id_log_produk" json:"log_produk"`
	Id_toko       uint           `gorm:"type:varchar(255);index" json:"id_toko"`
	Toko          Toko           `gorm:"foreignKey:Id_toko" json:"toko"`
	Kuantitas     int            `json:"kuantitas"`
	Harga_total   int            `json:"harga_total"`
	Created_at    time.Time      `json:"created_at"`
	Updated_at    time.Time      `json:"updated_at"`
	Deleted_at    gorm.DeletedAt `gorm:"index" json:"-"`
}
