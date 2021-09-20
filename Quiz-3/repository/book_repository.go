package repository

import (
	"Quiz-3/config"
	"Quiz-3/model"
	"context"
	"time"
)

func Insert(ctx context.Context, book model.Book) error {
	db := config.NewMySQL()
	
	_, err := db.ExecContext(ctx, "INSERT INTO book (title, description, image_url, release_year, price, total_page, kategori_ketebalan, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.KategoriKetebalan, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, book model.Book, bookId int) error {
	db := config.NewMySQL()
	
	_, err := db.ExecContext(ctx, "UPDATE book SET title = ?, description = ?, image_url = ?, release_year = ?, price = ?, total_page = ?, kategori_ketebalan = ?, updated_at = ? WHERE id = ?", book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.KategoriKetebalan, book.UpdatedAt, bookId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(ctx context.Context, bookId int) error {
	db := config.NewMySQL()
	
	_, err := db.ExecContext(ctx, "DELETE FROM book WHERE id = ?", bookId)
	if err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, query map[string]interface{}) ([]model.Book, error) {
	db := config.NewMySQL()

	// if query["title"] != "" {
	// 	if query["minYear"] != "" {
	// 		result, err := db.QueryContext(ctx, "SELECT * FROM book WHERE title = ? AND release_year > ?", query["title"], query["min_year"])
	// 	}
	// 	result, err := db.QueryContext(ctx, "SELECT * FROM book WHERE title = ?", query["title"])
	// }

	// if query["title"] != "" {
	// 	result, err := db.QueryContext(ctx, "SELECT * FROM book WHERE title = ?", query["title"])
	// }

	rows, err := db.QueryContext(ctx, "SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	
	var books []model.Book
	layoutDateTime := "2006-01-02 15:04:05"
	for rows.Next() {
		var book model.Book
		var createdAt, updatedAt string
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.KategoriKetebalan, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		book.CreatedAt, _ = time.Parse(layoutDateTime, createdAt)
		book.UpdatedAt, _ = time.Parse(layoutDateTime, updatedAt)

		books = append(books, book)
	}

	return books, nil
}