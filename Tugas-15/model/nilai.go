package model

import "time"

type NilaiMahasiswa struct{
	Nama string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai uint `json:"nilai"`
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostPutNilai struct {
	Nama string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	Nilai uint `json:"nilai"`
}