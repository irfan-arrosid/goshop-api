package requests

type TokoUpdateRequest struct {
	Nama_toko string `json:"nama_toko"`
}

// func (r *TokoUpdateRequest) MapRequest(url string) *entities.Toko {
// 	return &entities.Toko{
// 		Nama_toko: r.Nama_toko,
// 		Url_foto:  url,
// 	}
// }
