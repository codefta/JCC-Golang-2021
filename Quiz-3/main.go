package main

import (
	"Quiz-3/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/bangun-datar/segitiga-sama-sisi", controller.HitSegitigaSamaSisi)
	router.GET("/bangun-datar/Persegi", controller.HitPersegi)
	router.GET("/bangun-datar/persegi-panjang", controller.HitPersegiPanjang)
	router.GET("/bangun-datar/Lingkaran", controller.HitLingkaran)
	router.GET("/bangun-datar/jajar-genjang", controller.HitJajarGenjang)
	
	router.POST("/books", controller.AddBook)
	router.PUT("/books/:id", controller.EditBook)
	router.DELETE("/books/:id", controller.EditBook)
	router.GET("/books", controller.GetBook)
	m := Auth(router)

	fmt.Println("server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}

func Auth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "PUT" || r.Method == "DELETE" {
			uname, pwd, ok := r.BasicAuth()
			if !ok {
				http.Error(w, "Username atau Password tidak boleh kosong", http.StatusUnauthorized)
				return
			}
	
			if (uname == "admin" && pwd == "password") || (uname == "editor" && pwd == "secret") || (uname == "trainer" && pwd == "rahasia"){
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