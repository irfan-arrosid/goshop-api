package entities

import (
	"time"

	"gorm.io/gorm"
)

type Trx struct {
	Id                uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user           uint           `gorm:"type:varchar(255);index" json:"id_user"`
	User              User           `gorm:"foreignKey:Id_user" json:"user"`
	Alamat_pengiriman string         `gorm:"type:varchar(255)" json:"alamat_pengiriman"`
	Harga_total       int            `json:"harga_total"`
	Kode_invoice      string         `gorm:"type:varchar(255)" json:"kode_invoice"`
	Method_bayar      string         `gorm:"type:varchar(255)" json:"method_bayar"`
	Created_at        time.Time      `json:"created_at"`
	Updated_at        time.Time      `json:"updated_at"`
	Deleted_at        gorm.DeletedAt `gorm:"index" json:"-"`
}
