package requests

import "goshop-api/models/entities"

type (
	Trx struct {
		Method_bayar      string               `json:"method_bayar"`
		Alamat_pengiriman int                  `json:"alamat_pengiriman"`
		DetailTrx         []entities.DetailTrx `json:"detail_trx"`
	}

	DetailTrx struct {
		Id_log_produk uint `json:"Id_log_produk"`
		Kuantitas     uint `json:"kuantitas"`
	}
)
