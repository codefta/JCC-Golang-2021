package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

func UrutkanTampilkan(no int, p string, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	fmt.Printf("%d. %s\n", no, p)
}

func getMovies(channel chan string, movies ...string) {
	wg := &sync.WaitGroup{}
	no := 1
	for _, movie := range movies {
		wg.Add(1)
		channel <- fmt.Sprintf("%d. %s", no, movie)
		no++
		wg.Done()
	}
	wg.Wait()
	close(channel)
}

func LuasLingkaran(wg *sync.WaitGroup, channel chan float64, r int) {
	defer wg.Done()
	wg.Add(1)
	rFloat := float64(r)
	channel <- math.Pi * math.Pow(rFloat, 2)
}

func KelLingkaran(wg *sync.WaitGroup, channel chan float64, r int) {
	defer wg.Done()
	wg.Add(1)
	rFloat := float64(r)
	channel <- math.Pi * 2 * rFloat
}

func volTabung(wg *sync.WaitGroup, channel chan float64, r int, t int) {
	defer wg.Done()
	wg.Add(1)
	rFloat := float64(r)
	channel <- math.Pi * math.Pow(rFloat, 2) * float64(t)
}

func luasPersegiPanjang(wg *sync.WaitGroup, channel chan int, p int, l int) {
	defer wg.Done()
	wg.Add(1)
	channel <- p * l
}

func kelPersegiPanjang(wg *sync.WaitGroup, channel chan int, p, l int) {
	defer wg.Done()
	wg.Add(1)
	channel <- 2 * (p+l)
}

func volBalok(wg *sync.WaitGroup, channel chan int, p, l, t int) {
	defer wg.Done()
	wg.Add(1)
	channel <- p * l * t
}

func main() {
	// soal 1
	wg := &sync.WaitGroup{}
	phones := []string{
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo",
	}
	sort.Strings(phones)
	no := 1

	for _, p := range phones {
		go UrutkanTampilkan(no, p, wg)
		no++
		time.Sleep(1 * time.Second)
	}
	wg.Wait()

	// soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	wg = &sync.WaitGroup{}
	resultLuasLingkaran := make(chan float64)
	resultKelLingkaran := make(chan float64)
	resultVolTabung := make(chan float64)

	go LuasLingkaran(wg, resultLuasLingkaran, 8)
	fmt.Println(<- resultLuasLingkaran)
	go LuasLingkaran(wg, resultLuasLingkaran, 14)
	fmt.Println(<- resultLuasLingkaran)
	go LuasLingkaran(wg, resultLuasLingkaran, 20)
	fmt.Println(<- resultLuasLingkaran)
	close(resultLuasLingkaran)
	wg.Wait()

	go KelLingkaran(wg, resultKelLingkaran, 8)
	fmt.Println(<- resultKelLingkaran)
	go KelLingkaran(wg, resultKelLingkaran, 14)
	fmt.Println(<- resultKelLingkaran)
	go KelLingkaran(wg, resultKelLingkaran, 20)
	fmt.Println(<- resultKelLingkaran)
	close(resultKelLingkaran)
	wg.Wait()
	
	go volTabung(wg, resultVolTabung, 8, 10)
	fmt.Println(<- resultVolTabung)
	go volTabung(wg, resultVolTabung, 14, 10)
	fmt.Println(<- resultVolTabung)
	go volTabung(wg, resultVolTabung, 20, 10)
	fmt.Println(<- resultVolTabung)
	close(resultVolTabung)
	wg.Wait()

	// soal 4
	luasPPChan := make(chan int)
	kelPPChan := make(chan int)
	volBChan := make(chan int)
	defer close(luasPPChan)
	defer close(kelPPChan)
	defer close(volBChan)

	go luasPersegiPanjang(wg, luasPPChan, 7, 12)
	wg.Wait()
	go kelPersegiPanjang(wg, kelPPChan, 7, 12)
	wg.Wait()
	go volBalok(wg, volBChan, 7, 12, 5)
	wg.Wait()

	for i := 0; i < 3; i++ {
		select {
		case data := <- luasPPChan:
			fmt.Println("Luas Persegi Panjang", data)
		case data := <- kelPPChan:
			fmt.Println("Keliling Persegi Panjang", data)
		case data := <- volBChan:
			fmt.Println("Volume Balok", data)
		}
	}
}
