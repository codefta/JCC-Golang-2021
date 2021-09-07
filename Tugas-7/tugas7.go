package main

import "fmt"

type buah struct {
	nama, warna string
	adaBijinya bool
	harga int
}

type segitiga struct{
	alas, tinggi int
}
  
type persegi struct{
	sisi int
}

type persegiPanjang struct{
	panjang, lebar int
}

type phone struct{
	name, brand string
	year int
	colors []string
}

type movie struct{
	title, genre string
	duration, year int
}

func (s segitiga)lSegitiga() int {
	return (s.alas * s.tinggi) / 2
}

func (p persegi)lPersegi() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang)lPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

func (p *phone) addColor(newColor string) {
	p.colors = append(p.colors, newColor)
}

func tambahDataFilm(name string, duration int, genre string, year int, dataFilm *[]movie){
	*dataFilm = append(*dataFilm, movie{
		title: name,
		duration: duration,
		genre: genre,
		year: year,
	})
}

func main() {
	// soal 1
	nanas := buah{
		nama: "Nanas",
		warna: "Kuning",
		adaBijinya: false,
		harga: 9000,
	}
	fmt.Println(nanas)
	
	jeruk := buah{
		nama: "Jeruk",
		warna: "Oranye",
		adaBijinya: true,
		harga: 8000,
	}
	fmt.Println(jeruk)

	semangka := buah{
		nama: "Semangka",
		warna: "Hijau & Merah",
		adaBijinya: true,
		harga: 10000,
	}
	fmt.Println(semangka)

	pisang := buah{
		nama: "Pisang",
		warna: "Kuning",
		adaBijinya: false,
		harga: 5000,
	}
	fmt.Println(pisang)

	// soal 2
	luasSegitiga := segitiga{2, 5}
	fmt.Println(luasSegitiga.lSegitiga())

	luasPersegi := persegi{4}
	fmt.Println(luasPersegi.lPersegi())

	luasPersegiPanjang := persegiPanjang{5, 8}
	fmt.Println(luasPersegiPanjang.lPersegiPanjang())

	// soal 3
	xiaomi := phone{
		name: "Mi 11 Lite",
		brand: "Xiaomi",
		year: 2021,
		colors: []string{"Black"},
	}

	xiaomi.addColor("OceanBlue")
	fmt.Println(xiaomi)

	// soal 4
	dataFilm := []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	no := 1
	for _, item := range dataFilm {
		fmt.Printf("%d. title : %s\n   duration : %d\n   genre : %s\n   year : %d\n", no, item.title, item.duration, item.genre, item.year)
		no++
	}
}
