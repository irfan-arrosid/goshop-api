package models

type Province struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type City struct {
	Id         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Nama       string `json:"nama"`
}
