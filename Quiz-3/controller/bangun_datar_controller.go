package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func kelSegitigaSamaSisi(wg *sync.WaitGroup, channel chan int, a int) {
	defer wg.Done()
	wg.Add(1)
	channel <- 3 * a
}

func luasSegitigaSamaSisi(wg *sync.WaitGroup, channel chan int, a int, t int) {
	defer wg.Done()
	wg.Add(1)
	channel <- (a * t) / 2
}

func luasPersegi(wg *sync.WaitGroup, channel chan int, s int) {
	defer wg.Done()
	wg.Add(1)
	channel <- s * s
}

func kelPersegi(wg *sync.WaitGroup, channel chan int, s int) {
	defer wg.Done()
	wg.Add(1)
	channel <- 4 * s
}

func kelPersegiPanjang(wg *sync.WaitGroup, channel chan int, p int, l int) {
	defer wg.Done()
	wg.Add(1)
	channel <- 2 * (p + l)
}

func luasPersegiPanjang(wg *sync.WaitGroup, channel chan int, p int, l int) {
	defer wg.Done()
	wg.Add(1)
	channel <- p * l
}

func kelLingkaran(wg *sync.WaitGroup, channel chan float64, r int) {
	defer wg.Done()
	wg.Add(1)
	rFloat := float64(r)
	channel <- math.Pi * 2 * rFloat
}

func luasLingkaran(wg *sync.WaitGroup, channel chan float64, r int) {
	defer wg.Done()
	wg.Add(1)
	rFloat := float64(r)
	channel <- math.Pi * math.Pow(rFloat, 2)
}

func kelJajarGenjang(wg *sync.WaitGroup, channel chan int, a int, s int) {
	defer wg.Done()
	wg.Add(1)
	channel <- 2 * a + 2 * s
}

func luasJajarGenjang(wg *sync.WaitGroup, channel chan int, a int, t int) {
	defer wg.Done()
	wg.Add(1)
	channel <- a * t
}

func HitSegitigaSamaSisi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	wg := &sync.WaitGroup{}

	hitung := queryValues.Get("hitung")
	if hitung != "" {
		a := queryValues.Get("alas")
		alas, _ := strconv.Atoi(a)
		t := queryValues.Get("tinggi")
		tinggi, _ := strconv.Atoi(t)

		if hitung == "luas" {
			lSS := make(chan int)
			go luasSegitigaSamaSisi(wg, lSS, alas, tinggi)
			wg.Wait()
			l := <- lSS
			close(lSS)
			fmt.Fprintf(w, "Luas Segitiga Sama Sisi: %d", l)
			return
		} else if hitung == "keliling" {
			result := make(chan int)
			go kelSegitigaSamaSisi(wg, result, alas)
			wg.Wait()
			l := <- result
			close(result)
			fmt.Fprintf(w, "Keliling Segitiga Sama Sisi: %d", l)
			return
		}
	} else {
		http.Error(w, "Anda tidak memasukan query params", 400)
		return
	}
}

func HitPersegi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	wg := &sync.WaitGroup{}

	hitung := queryValues.Get("hitung")
	if hitung != "" {
		s := queryValues.Get("sisi")
		sisi, _ := strconv.Atoi(s)

		if hitung == "luas" {
			lSS := make(chan int)
			go luasPersegi(wg, lSS, sisi)
			wg.Wait()
			l := <- lSS
			close(lSS)
			fmt.Fprintf(w, "Luas Persegi: %d", l)
			return
		} else if hitung == "keliling" {
			result := make(chan int)
			go kelPersegi(wg, result, sisi)
			wg.Wait()
			l := <- result
			close(result)
			fmt.Fprintf(w, "Keliling Persegi: %d", l)
			return
		}
	} else {
		http.Error(w, "Anda tidak memasukan query params", 400)
		return
	}
}

func HitLingkaran(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	wg := &sync.WaitGroup{}

	hitung := queryValues.Get("hitung")
	if hitung != "" {
		l := queryValues.Get("jariJari")
		jariJari, _ := strconv.Atoi(l)

		if hitung == "luas" {
			lSS := make(chan float64)
			go luasLingkaran(wg, lSS, jariJari)
			wg.Wait()
			l := <- lSS
			close(lSS)
			fmt.Fprintf(w, "Luas Lingkaran: %f", l)
			return
		} else if hitung == "keliling" {
			result := make(chan float64)
			go kelLingkaran(wg, result, jariJari)
			wg.Wait()
			l := <- result
			close(result)
			fmt.Fprintf(w, "Keliling Lingkaran: %f", l)
			return
		}
	} else {
		http.Error(w, "Anda tidak memasukan query params", 400)
		return
	}
}

func HitPersegiPanjang(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	wg := &sync.WaitGroup{}

	hitung := queryValues.Get("hitung")
	if hitung != "" {
		p := queryValues.Get("panjang")
		panjang, _ := strconv.Atoi(p)
		l := queryValues.Get("lebar")
		lebar, _ := strconv.Atoi(l)

		if hitung == "luas" {
			lSS := make(chan int)
			go luasPersegiPanjang(wg, lSS, panjang, lebar)
			wg.Wait()
			l := <- lSS
			close(lSS)
			fmt.Fprintf(w, "Luas Persegi Panjang: %d", l)
			return
		} else if hitung == "keliling" {
			result := make(chan int)
			go kelPersegiPanjang(wg, result, panjang, lebar)
			wg.Wait()
			l := <- result
			close(result)
			fmt.Fprintf(w, "Keliling Persegi Panjang: %d", l)
			return
		}
	} else {
		http.Error(w, "Anda tidak memasukan query params", 400)
		return
	}
}

func HitJajarGenjang(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	wg := &sync.WaitGroup{}

	hitung := queryValues.Get("hitung")
	if hitung != "" {
		a := queryValues.Get("alas")
		alas, _ := strconv.Atoi(a)
		t := queryValues.Get("tinggi")
		tinggi, _ := strconv.Atoi(t)
		s := queryValues.Get("sisi")
		sisi, _ := strconv.Atoi(s)

		if hitung == "luas" {
			lSS := make(chan int)
			go luasJajarGenjang(wg, lSS, alas, tinggi)
			wg.Wait()
			l := <- lSS
			close(lSS)
			fmt.Fprintf(w, "Luas Jajar Genjang: %d", l)
			return
		} else if hitung == "keliling" {
			result := make(chan int)
			go kelJajarGenjang(wg, result, alas, sisi)
			wg.Wait()
			l := <- result
			close(result)
			fmt.Fprintf(w, "Keliling Jajar Genjang: %d", l)
			return
		}
	} else {
		http.Error(w, "Anda tidak memasukan query params", 400)
		return
	}
}