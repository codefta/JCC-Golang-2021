package main

import (
	"fmt"
	"math"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (ss segitigaSamaSisi) luas() int {
	return ss.alas * ss.tinggi / 2
}

func (ss segitigaSamaSisi) keliling() int {
	return 3 * ss.alas
}

func (pp persegiPanjang) keliling() int {
	return 2 * (pp.panjang + pp.lebar)
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)	
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64(b.panjang * b.lebar + b.panjang * b.tinggi + b.lebar * b.tinggi)
}

type phone struct {
	name, brand string
	year int
	colors []string
}

type phoneSpec interface {
	showPhoneSpec() string
}

func (p phone) showPhoneSpec() string {
	return fmt.Sprintf("name : %v\nbrand : %v\nyear : %d\ncolors : %v", p.name, p.brand, p.year, strings.TrimRight(strings.Join(p.colors, ", "), ", "))
}

func luasPersegi(sisi int, status bool) interface{} {
	if sisi > 0 && status == true {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, sisi * sisi)
	} else if sisi > 0 && status == false {
		return sisi * sisi
	} else if sisi == 0 && status == true {
		return "Maaf anda belum menginput sisi dari persegi"
	} else {
		return nil
	}
}

func main() {
	// soal 1
	var bangunDatar hitungBangunDatar

	bangunDatar = segitigaSamaSisi{5, 5}
	fmt.Println("Luas segitiga sama sisi", bangunDatar.luas())
	fmt.Println("Keliling segitiga sama sisi", bangunDatar.keliling())

	bangunDatar = persegiPanjang{3, 8}
	fmt.Println("Luas persegi panjang", bangunDatar.luas())
	fmt.Println("Keliling persegi panjang", bangunDatar.keliling())

	var bangunRuang hitungBangunRuang
	bangunRuang = tabung{10, 20}
	fmt.Println("Volume tabung", bangunRuang.volume())
	fmt.Println("Luas Permukaan tabung", bangunRuang.luasPermukaan())

	bangunRuang = balok{5, 10, 5}
	fmt.Println("Volume balok", bangunRuang.volume())
	fmt.Println("Luas Permukaan balok", bangunRuang.luasPermukaan())

	// soal 2
	var myPhone phoneSpec
	
	myPhone = phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(myPhone.showPhoneSpec())

	// soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6,8}

	var kumpulanAngkaKedua interface{} = []int{12,14}

	prefixStr := prefix.(string)
	angkaPertamaIntArr := kumpulanAngkaPertama.([]int)
	angkaKeduaIntArr := kumpulanAngkaKedua.([]int)

	sumAngka := 0
	for _, angka := range angkaPertamaIntArr {
		sumAngka += angka
	}

	for _, angka := range angkaKeduaIntArr {
		sumAngka += angka
	}

	result := fmt.Sprintf("%v%v + %v = %d", prefixStr, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(angkaPertamaIntArr)), " + "), "[]"), strings.Trim(strings.TrimRight(strings.Join(strings.Fields(fmt.Sprint(angkaKeduaIntArr)), " + "), " + "), "[]"), sumAngka)
	
	fmt.Println(result)
}
