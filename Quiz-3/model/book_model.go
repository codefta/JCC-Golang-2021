package model

import "time"

type Book struct {
	ID 					int			`json:"id"`
	Title				string		`json:"title"`
	Description			string		`json:"description"`
	ImageURL			string		`json:"image_url"`
	ReleaseYear			int			`json:"release_year"`
	Price				string		`json:"price"`
	TotalPage			string		`json:"total_page"`
	KategoriKetebalan	string		`json:"kategori_ketebalan"`
	CreatedAt			time.Time	`json:"created_at"`
	UpdatedAt			time.Time	`json:"updated_at"`
}

type PostPutBook struct {
	Title				string		`json:"title"`
	Description			string		`json:"description"`
	ImageURL			string		`json:"image_url"`
	ReleaseYear			int			`json:"release_year"`
	Price				int			`json:"price"`
	TotalPage			string		`json:"total_page"`
	CreatedAt			time.Time	`json:"created_at"`
	UpdatedAt			time.Time	`json:"updated_at"`
}