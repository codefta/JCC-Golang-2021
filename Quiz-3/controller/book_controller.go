package controller

import (
	"Quiz-3/helper"
	"Quiz-3/model"
	"Quiz-3/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bookPost model.PostPutBook
	err := json.NewDecoder(r.Body).Decode(&bookPost)
	if err != nil {
		helper.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var validation = []string{}

	if bookPost.ReleaseYear < 1980 || bookPost.ReleaseYear > 2021 {
		validation = append(validation, "Inputan harus lebih dari sama dengan 1980 dan tidak boleh lebih daripada 2021")
	}

	_, err = url.ParseRequestURI(bookPost.ImageURL)
	if err != nil {
		validation = append(validation, "Image URL tidak valid")
	}

	if len(validation) > 0 {
		helper.ResponseJSON(w, validation, http.StatusBadRequest)
		return
	}
	
	var book model.Book
	book.Title = bookPost.Title
	book.Description = bookPost.Description
	book.ImageURL = bookPost.ImageURL
	book.TotalPage = bookPost.TotalPage
	book.ReleaseYear = bookPost.ReleaseYear
	book.CreatedAt = bookPost.CreatedAt
	book.UpdatedAt = bookPost.UpdatedAt
	toRupiah := fmt.Sprintf("Rp. %d", bookPost.Price)
	book.Price = toRupiah
	totalPage, _ := strconv.Atoi(bookPost.TotalPage)
	if  totalPage <= 100 {
		book.TotalPage = "tipis"
	} else if totalPage >= 101 && totalPage <= 200 {
		book.TotalPage = "sedang"
	} else if totalPage >= 201 {
		book.TotalPage = "tebal"
	}

	err = repository.Insert(ctx, book)
	if err != nil {
		helper.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "Data berhasil ditambahkan",
	}

	helper.ResponseJSON(w, res , http.StatusOK)
}

func EditBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := p.ByName("id")
	bookId, _ := strconv.Atoi(id)
	var bookPut model.PostPutBook
	err := json.NewDecoder(r.Body).Decode(&bookPut)
	if err != nil {
		helper.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var validation = []string{}

	if bookPut.ReleaseYear < 1980 || bookPut.ReleaseYear > 2021 {
		validation = append(validation, "Inputan harus lebih dari sama dengan 1980 dan tidak boleh lebih daripada 2021")
	}

	_, err = url.ParseRequestURI(bookPut.ImageURL)
	if err != nil {
		validation = append(validation, "Image URL tidak valid")
	}

	if len(validation) > 0 {
		helper.ResponseJSON(w, validation, http.StatusBadRequest)
		return
	}
	
	var book model.Book
	book.Title = bookPut.Title
	book.Description = bookPut.Description
	book.ImageURL = bookPut.ImageURL
	book.TotalPage = bookPut.TotalPage
	book.ReleaseYear = bookPut.ReleaseYear
	book.UpdatedAt = bookPut.UpdatedAt
	toRupiah := fmt.Sprintf("Rp. %d", bookPut.Price)
	book.Price = toRupiah
	totalPage, _ := strconv.Atoi(bookPut.TotalPage)
	if  totalPage <= 100 {
		book.TotalPage = "tipis"
	} else if totalPage >= 101 && totalPage <= 200 {
		book.TotalPage = "sedang"
	} else if totalPage >= 201 {
		book.TotalPage = "tebal"
	}

	err = repository.Update(ctx, book, bookId)
	if err != nil {
		helper.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "Data berhasil diupdate",
	}

	helper.ResponseJSON(w, res , http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := p.ByName("id")
	bookId, _ := strconv.Atoi(id)

	err := repository.Delete(ctx, bookId)
	if err != nil {
		helper.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "Data berhasil dihapus",
	}

	helper.ResponseJSON(w, res , http.StatusOK)
}

func GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	queryValues := r.URL.Query()
	title := queryValues.Get("title")
	sort := queryValues.Get("sort")
	miPage := queryValues.Get("minPage")
	minPage, _ := strconv.Atoi(miPage)
	maPage := queryValues.Get("maxPage")
	maxPage, _ := strconv.Atoi(maPage)
	maYear := queryValues.Get("maxYear")
	maxYear, _ := strconv.Atoi(maYear)
	miYear := queryValues.Get("minYear")
	minYear, _ := strconv.Atoi(miYear)

	query := map[string]interface{}{
		"title": title,
		"minYear": minYear,
		"maxYear": maxYear,
		"minPage": minPage,
		"maxPage": maxPage,
		"sort": sort,
	}

	result, err := repository.Get(ctx, query)
	if err != nil {
		helper.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	helper.ResponseJSON(w, result , http.StatusOK)
}