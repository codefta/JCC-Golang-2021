package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "jcc"
	password string = "Jcc@2021"
	database string = "books"
)

func NewMySQL() *sql.DB {
	stmt := fmt.Sprintf("%v:%v@/%v", username, password, database)
	db, err := sql.Open("mysql", stmt)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}