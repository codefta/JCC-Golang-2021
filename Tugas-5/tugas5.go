package main

import (
	"fmt"
	"strings"
)

func luasPersegiPanjang(p int, l int) int {
	return p * l
}

func kelilingPersegiPanjang(p int, l int) int {
	return 2 * (p + l)
}

func volumeBalok(p int, l int, t int) int {
	return p * l * t
}

func introduce(name, gender, job, age string) string {
	return fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
}

func buahFavorit(name string, fruits ...string) string {
	resultString := fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah ", strings.ToLower(name))
	for _, fruit := range fruits {
		resultString = resultString + fmt.Sprintf("\"%s\", ", fruit)
	}
	return strings.TrimRight(resultString, ", ")
}

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8
  
	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas) 
	fmt.Println(keliling)
	fmt.Println(volume)
	
	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	
	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
	
	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(title, jam, genre, tahun string){
		dataFilm = append(dataFilm, map[string]string{
			"genre": genre,
			"jam": jam,
			"tahun": tahun,
			"title": title,
		})
	} 

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
