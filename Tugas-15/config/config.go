package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "jcc"
	password string = "Jcc@2021"
	database string = "nilai_mhs"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db, nil
}