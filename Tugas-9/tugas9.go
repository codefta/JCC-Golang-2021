package main

import (
	"errors"
	"fmt"
	"strconv"
)

func jcc(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

func runJcc() {
	defer jcc("Candradimuka Jabar Coding Camp", 2021)
	fmt.Println("JCC Terjalankan")
}

func kelilingSegitigaSamaSisi(sisi int, status bool) (string, error) {
	defer reco()
	if (sisi > 0 && status == true) {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, sisi*sisi), nil
	} else if (sisi > 0 && status == false) {
		return strconv.Itoa(sisi), nil
	} else if (sisi == 0 && status == true) {
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if (sisi == 0 && status == false) {
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	return "", nil
}

func reco() {
	if message := recover(); message != nil {
		fmt.Println("Terjadi panic: ", message)
	}
}

func tambahAngka(increment int, result *int) {
	*result = *result + increment
}

func main() {
	angka := 1
	defer fmt.Println(angka)

	// soal 1
	runJcc()

	// soal 2
	result, _ := kelilingSegitigaSamaSisi(4, true)
	fmt.Println("Luas Segitiga: ", result)

	result2, _ := kelilingSegitigaSamaSisi(8, false)
	fmt.Println("Luas Segitiga: ", result2)

	_, err := kelilingSegitigaSamaSisi(0, true)
	if err != nil {
		fmt.Println(err)
	}

	kelilingSegitigaSamaSisi(0, false)

	// soal 3
	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)
}
