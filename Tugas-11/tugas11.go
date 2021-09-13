package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

func UrutkanTampilkan(phones *[]string) {
	*phones = []string{
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo",
	}

	sort.Strings(*phones)

	no := 1
	for _, p := range *phones {
		fmt.Printf("%d. %s\n", no, p)
		no++
		time.Sleep(1 * time.Second)
	}
}

func LuasLingkaran(r int) float64{
	rFloat := float64(r)
	return math.Round(math.Pi * math.Pow(rFloat, 2))
}

func KelLingkaran(r int) float64 {
	rFloat := float64(r)
	return math.Round(math.Pi * 2 * rFloat)
}

func main() {
	// soal 1
	phones := []string{}
	UrutkanTampilkan(&phones)

	// soal 2
	fmt.Println(LuasLingkaran(7))
	fmt.Println(LuasLingkaran(10))
	fmt.Println(LuasLingkaran(15))

	fmt.Println(KelLingkaran(7))
	fmt.Println(KelLingkaran(10))
	fmt.Println(KelLingkaran(15))

	// soal 3
	p := flag.Int("Panjang", 5, "Panjang Persegi Panjang")
	l := flag.Int("Lebar", 10, "Lebar Persegi Panjang")
	flag.Parse()
	
	luas := *p * *l
	keliling := 2 * (*p + *l)
	fmt.Println(luas)
	fmt.Println(keliling)
}
