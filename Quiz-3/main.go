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

	fmt.Println("server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}