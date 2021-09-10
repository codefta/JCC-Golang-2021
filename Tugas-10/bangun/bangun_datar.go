package bangun

import "fmt"

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

func (ss SegitigaSamaSisi) Luas() int {
	return ss.Alas * ss.Tinggi / 2
}

func (ss SegitigaSamaSisi) Keliling() int {
	return 3 * ss.Alas
}

func (pp PersegiPanjang) Keliling() int {
	return 2 * (pp.Panjang + pp.Lebar)
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func LuasPersegi(sisi int, status bool) interface{} {
	if sisi > 0 && status {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, sisi * sisi)
	} else if sisi > 0 && !status {
		return sisi * sisi
	} else if sisi == 0 && status {
		return "Maaf anda belum menginput sisi dari persegi"
	} else {
		return nil
	}
}