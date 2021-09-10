package main

import (
	"Tugas-10/bangun"
	"Tugas-10/tech"
	"fmt"
	"strings"
)

func main() {
	// soal 1
	var bangunDatar bangun.HitungBangunDatar

	bangunDatar = bangun.SegitigaSamaSisi{5, 5}
	fmt.Println("Luas segitiga sama sisi", bangunDatar.Luas())
	fmt.Println("Keliling segitiga sama sisi", bangunDatar.Keliling())

	bangunDatar = bangun.PersegiPanjang{3, 8}
	fmt.Println("Luas persegi panjang", bangunDatar.Luas())
	fmt.Println("Keliling persegi panjang", bangunDatar.Keliling())

	var bangunRuang bangun.HitungBangunRuang
	bangunRuang = bangun.Tabung{10, 20}
	fmt.Println("Volume tabung", bangunRuang.Volume())
	fmt.Println("Luas Permukaan tabung", bangunRuang.LuasPermukaan())

	bangunRuang = bangun.Balok{5, 10, 5}
	fmt.Println("Volume balok", bangunRuang.Volume())
	fmt.Println("Luas Permukaan balok", bangunRuang.LuasPermukaan())

	// soal 2
	var myPhone tech.PhoneSpec
	
	myPhone = tech.Phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(myPhone.ShowPhoneSpec())

	// soal 3
	fmt.Println(bangun.LuasPersegi(4, true))

	fmt.Println(bangun.LuasPersegi(8, false))

	fmt.Println(bangun.LuasPersegi(0, true))

	fmt.Println(bangun.LuasPersegi(0, false))

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