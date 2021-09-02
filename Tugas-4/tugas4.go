package main

import (
	"fmt"
)

func main() {
	// soal 1
	n := 20
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d - Candradimuka\n", i)
		} else {
			fmt.Printf("%d - JCC\n", i)
		}
	}

	// soal 2
	for i := 0; i < 7; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// soal 3
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	fmt.Println(kalimat[2:6])

	// soal 4
	sayuran := []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	no := 1
	for i := 0; i < len(sayuran); i++ {
		fmt.Printf("%d. %s\n", no, sayuran[i])
		no++
	}

	// soal 5
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	
	v := 1
	for i, s := range satuan {
		fmt.Println(i, s)
		v = v*s
	}
	
	fmt.Printf("Volume Balok = %d", v)

}
