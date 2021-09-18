package main

import (
	"Tugas-15/model"
	"Tugas-15/repository"
	"Tugas-15/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nilai", getNilai)
	router.GET("/nilai/:id", getNilaiId)
	router.POST("/nilai", postNilai)
	router.PUT("/nilai/:id", putNilai)
	router.DELETE("/nilai/:id", deleteNilai)
	m := Auth(router)

	fmt.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}

func getNilai(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nilaiMhs, err := repository.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilaiMhs, http.StatusOK)
}

func getNilaiId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idMhs := p.ByName("id")
	id, _ := strconv.Atoi(idMhs)

	nilaiMhs, err := repository.FindById(ctx, id)
	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ResponseJSON(w, nilaiMhs, http.StatusOK)
}

func postNilai(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var postNilai model.PostPutNilai
	err := json.NewDecoder(r.Body).Decode(&postNilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	err = repository.Insert(ctx, postNilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	res := map[string]string{
		"status": "Success insert data",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func putNilai(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var putNilai model.PostPutNilai
	err := json.NewDecoder(r.Body).Decode(&putNilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idMhs := p.ByName("id")
	id, _ := strconv.Atoi(idMhs)

	err = repository.Update(ctx, putNilai, id)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "Update successfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func deleteNilai(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idMhs := p.ByName("id")
	id, _ := strconv.Atoi(idMhs)

	err := repository.Delete(ctx, id)
	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "Delete successfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" {
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