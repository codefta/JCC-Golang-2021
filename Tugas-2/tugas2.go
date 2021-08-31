package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	jabar := "Jabar"
	coding := "Coding"
	camp := "Camp"
	subjekPelatihan := "Golang"
	tahun := 2021
	fmt.Println(jabar, coding, camp, subjekPelatihan, tahun)

	// soal 2
	halo := "Halo Dunia"
	golang := strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println(golang)

	// soal 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";
	fmt.Println(kataPertama, strings.Replace(kataKedua, "s", "S", 1), strings.Replace(kataKetiga, "r", "R", 1), strings.ToUpper(kataKeempat))

	// soal 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";

	intPertama, _ := strconv.Atoi(angkaPertama)
	intKedua, _ := strconv.Atoi(angkaKedua)
	intKetiga, _ := strconv.Atoi(angkaKetiga)
	intKeempat, _ := strconv.Atoi(angkaKeempat)
	total := intPertama + intKedua + intKetiga + intKeempat
	fmt.Println(total)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	fmt.Printf("\"%s\"-%d", strings.ReplaceAll(kalimat, "halo", "Hi"), angka)
}