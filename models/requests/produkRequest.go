package requests

import (
	"goshop-api/models/entities"
	"strings"
)

type ProdukCreateRequest struct {
	Nama_produk    string `gorm:"type:varchar(255);index" json:"nama_produk"`
	Harga_reseller string `gorm:"type:varchar(255)" json:"harga_reseller"`
	Harga_konsumen string `gorm:"type:varchar(255)" json:"harga_konsumen"`
	Stok           int    `json:"stok"`
	Deskripsi      string `json:"deskripsi"`
	Id_category    uint   `gorm:"type:varchar(255);index" json:"id_category"`
}

func (r *ProdukCreateRequest) MapRequest(foto []entities.FotoProduk) *entities.Produk {
	slug := strings.ToLower(strings.ReplaceAll(r.Nama_produk, " ", "-"))
	return &entities.Produk{
		Nama_produk:    r.Nama_produk,
		Slug:           slug,
		Id_category:    r.Id_category,
		Harga_reseller: r.Harga_reseller,
		Harga_konsumen: r.Harga_konsumen,
		Stok:           r.Stok,
		Deskripsi:      r.Deskripsi,
		Foto:           foto,
	}
}
