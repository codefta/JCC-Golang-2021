package main

import (
	"fmt"
	"math"
)

func luasL(luas *float64, r float64) {
	*luas = math.Phi * float64(math.Abs(r))
}

func kelL(kel *float64, r float64) {
	*kel = 2.00 * math.Phi * r
}

func introduce(sentence *string, name, gender, job, age string) {
	if gender == "laki-laki" {
		*sentence = "Pak " + name + " adalah seorang "  + job + " yang berusia " + age
	} else {
		*sentence = "Bu " + name + " adalah seorang "  + job + " yang berusia " + age
	}
}

func addBuah(buah *[]string, buahNew string) {
	*buah = append(*buah, buahNew)
}


func main() {
	// soal 1
	var luasLigkaran float64 
	var kelilingLingkaran float64

	luasL(&luasLigkaran, 8.00)
	kelL(&kelilingLingkaran, 8.00)

	// soal 2
	var sentence string 

	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{}
	addBuah(&buah, "Jeruk")
	addBuah(&buah, "Semangka")
	addBuah(&buah, "Mangga")
	addBuah(&buah, "Strawberry")
	addBuah(&buah, "Durian")
	addBuah(&buah, "Manggis")
	addBuah(&buah, "Alpukat")

	no := 1
	for _, b := range buah {
		fmt.Printf("%d.%s\n", no, b)
		no++
	}

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(title, jam, genre, tahun string, dataFilm *[]map[string]string){
		*dataFilm = append(*dataFilm, map[string]string{
			"genre": genre,
			"jam": jam,
			"tahun": tahun,
			"title": title,
		})
	} 

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	
	no = 1
	for _, item := range dataFilm {
		fmt.Printf("%d. title : %s\nduration : %s\ngenre : %s\nyear : %s\n", no, item["title"], item["jam"], item["genre"], item["tahun"])
	}
}
