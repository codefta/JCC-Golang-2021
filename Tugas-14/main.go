package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type NilaiMahasiswa struct{
	Nama string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai uint `json:"nilai"`
	ID uint `json:"id"`
}

type PostNilai struct {
	Nama string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	Nilai uint `json:"nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{

}

func getPostNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nilai := &nilaiNilaiMahasiswa
		dataNilai, err := json.Marshal(*nilai)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	} else if r.Method == "POST" {
		postNilai := PostNilai{}
		nilaiResult := NilaiMahasiswa{}

		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&postNilai); err != nil {
				log.Fatal(err)
			}
		} else {
			http.Error(w, "Data must be json format", http.StatusBadRequest)
			return
		}

		nilaiResult.Nama = postNilai.Nama
		nilaiResult.MataKuliah = postNilai.MataKuliah
		nilaiResult.Nilai = postNilai.Nilai

		if len(nilaiNilaiMahasiswa) > 0 {
			nilaiResult.ID = nilaiNilaiMahasiswa[len(nilaiNilaiMahasiswa) - 1].ID + 1
		} else {
			nilaiResult.ID = 1
		}

		if postNilai.Nilai > 100 {
			http.Error(w, "Nilai tidak bisa lebih dari 100", http.StatusBadRequest)
			return
		}

		if nilaiResult.Nilai >= 80 {
			nilaiResult.IndeksNilai = "A"
		} else if nilaiResult.Nilai >= 70 && nilaiResult.Nilai < 80 {
			nilaiResult.IndeksNilai = "B"
		} else if nilaiResult.Nilai >= 60 && nilaiResult.Nilai < 70 {
			nilaiResult.IndeksNilai = "C"
		} else if nilaiResult.Nilai >= 50 && nilaiResult.Nilai < 60 {
			nilaiResult.IndeksNilai = "D"
		} else if nilaiResult.Nilai < 50 {
			nilaiResult.IndeksNilai = "E"
		}
		
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiResult)
		nilaiPosted, err:= json.Marshal(nilaiResult)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nilaiPosted)
		return
	}
	http.Error(w, "ERRR>>>>", http.StatusNotFound)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			uname, pwd, ok := r.BasicAuth()
			if !ok {
				http.Error(w, "Username atau Password tidak boleh kosong", http.StatusUnauthorized)
				return
			}
	
			if uname == "admin" && pwd == "admin" {
				next.ServeHTTP(w, r)
				return
			}
	
			http.Error(w, "Username atau Password tidak salah", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
				return
		}
	})
}

func main() {
	http.Handle("/nilai", Auth(http.HandlerFunc(getPostNilai)))

	err := http.ListenAndServe(":8080", nil); 
	if err != nil {
		log.Fatal(err)
	}
}