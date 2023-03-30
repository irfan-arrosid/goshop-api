package requests

type AlamatCreateRequest struct {
	Judul_alamat  string `json:"judul_alamat"`
	Nama_penerima string `json:"nama_penerima"`
	Notelp        string `json:"notelp"`
	Detail_alamat string `json:"detail_alamat"`
}
