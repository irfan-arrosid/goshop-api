package requests

import (
	"time"
)

type UserCreateRequest struct {
	Nama          string    `json:"nama" validate:"required"`
	Kata_sandi    string    `json:"kata_sandi" validate:"required"`
	Notelp        string    `json:"notelp"`
	Tanggal_lahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	Jenis_kelamin string    `json:"jenis_kelamin"`
	Tentang       string    `json:"tentang"`
	Pekerjaan     string    `json:"pekerjaan"`
	Email         string    `json:"email" validate:"required"`
	Id_provinsi   string    `json:"id_provinsi"`
	Id_kota       string    `json:"id_kota"`
	IsAdmin       bool      `json:"isAdmin"`
}

type UserUpdateRequest struct {
	Nama          string    `json:"nama" validate:"required"`
	Notelp        string    `json:"notelp"`
	Tanggal_lahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	Jenis_kelamin string    `json:"jenis_kelamin"`
	Tentang       string    `json:"tentang"`
	Pekerjaan     string    `json:"pekerjaan"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type UserLoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Kata_sandi string `json:"kata_sandi" validate:"required"`
}
