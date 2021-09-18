package repository

import (
	"Tugas-15/config"
	"Tugas-15/model"
	"context"
	"fmt"
	"log"
	"time"
)

const (
	table = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]model.NilaiMahasiswa, error) {
	var nilaiMhs []model.NilaiMahasiswa
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v", table)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var nilai model.NilaiMahasiswa
		var createdAt, updatedAt string
		err = rows.Scan(&nilai.ID, &nilai.Nama, &nilai.MataKuliah, &nilai.Nilai, &nilai.IndeksNilai, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		nilaiMhs = append(nilaiMhs, nilai)
	}

	return nilaiMhs, nil
}

func Insert(ctx context.Context, postNilai model.PostPutNilai) error {
	var nilai model.NilaiMahasiswa
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	nilai.Nama = postNilai.Nama
	nilai.MataKuliah = postNilai.MataKuliah
	nilai.Nilai = postNilai.Nilai

	if nilai.Nilai >= 80 {
		nilai.IndeksNilai = "A"
	} else if nilai.Nilai >= 70 && nilai.Nilai < 80 {
		nilai.IndeksNilai = "B"
	} else if nilai.Nilai >= 60 && nilai.Nilai < 70 {
		nilai.IndeksNilai = "C"
	} else if nilai.Nilai >= 50 && nilai.Nilai < 60 {
		nilai.IndeksNilai = "D"
	} else if nilai.Nilai < 50 {
		nilai.IndeksNilai = "E"
	}


	queryText := "INSERT INTO nilai (nama, mata_kuliah, nilai, indeks_nilai, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())"

	_, err = db.ExecContext(ctx, queryText, nilai.Nama, nilai.MataKuliah, nilai.Nilai, nilai.IndeksNilai)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Update(ctx context.Context, putNilai model.PostPutNilai, idMhs int) error {
	var nilai model.NilaiMahasiswa
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	nilai.Nama = putNilai.Nama
	nilai.MataKuliah = putNilai.MataKuliah
	nilai.Nilai = putNilai.Nilai

	if nilai.Nilai >= 80 {
		nilai.IndeksNilai = "A"
	} else if nilai.Nilai >= 70 && nilai.Nilai < 80 {
		nilai.IndeksNilai = "B"
	} else if nilai.Nilai >= 60 && nilai.Nilai < 70 {
		nilai.IndeksNilai = "C"
	} else if nilai.Nilai >= 50 && nilai.Nilai < 60 {
		nilai.IndeksNilai = "D"
	} else if nilai.Nilai < 50 {
		nilai.IndeksNilai = "E"
	}

	queryText := "UPDATE nilai set nama = ? ,mata_kuliah = ?, nilai = ?, indeks_nilai = ?, updated_at = NOW() WHERE id = ?"

	_, err = db.ExecContext(ctx, queryText, nilai.Nama, nilai.MataKuliah, nilai.Nilai, nilai.IndeksNilai, idMhs)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Delete(ctx context.Context, idMhs int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	_, err = FindById(ctx, idMhs)
	if err != nil {
		return err
	}

	queryText := "DELETE FROM nilai WHERE id = ?"

	_, err = db.ExecContext(ctx, queryText, idMhs)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func FindById(ctx context.Context, idMhs int) (*model.NilaiMahasiswa, error) {
	var nilaiMhs model.NilaiMahasiswa
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v WHERE id = %v", table, idMhs)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {
		var createdAt, updatedAt string
		err = rows.Scan(&nilaiMhs.ID, &nilaiMhs.Nama, &nilaiMhs.MataKuliah, &nilaiMhs.Nilai, &nilaiMhs.IndeksNilai, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		nilaiMhs.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		nilaiMhs.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil, fmt.Errorf("id tidak ditemukan")
	}

	return &nilaiMhs, nil
}