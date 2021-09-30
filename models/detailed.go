package models

import (
	"fmt"
	"go-postgres-crud/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Buku schema dari tabel Buku
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Detailed struct {
	Id       int64 `json:"id"`
	Season   int64 `json:"season"`
	Episodes int64 `json:"episodes"`
	Year     int64 `json:"year"`
}

func AddDetailed(detail Detailed) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari buku yang dimasukkan ke db
	sqlStatement := `INSERT INTO detailed (id, season, episodes, year) VALUES ($1, $2, $3, $4) RETURNING id`

	// id yang dimasukkan akan disimpan di id ini
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, detail.Id, detail.Season, detail.Episodes, detail.Year).Scan(&id)

	if err != nil {
		log.Fatalf("Query could not be executed. %v", err)
	}

	fmt.Printf("Data inserted in record %v", id)

	// return insert id
	return id
}

// ambil buku
func GetDetaildata() ([]Detailed, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var details []Detailed

	// kita buat select query
	sqlStatement := `SELECT * FROM detailed`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Query could not be executed. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var detail Detailed

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&detail.Id, &detail.Season, &detail.Episodes, &detail.Year)

		if err != nil {
			log.Fatalf("No data. %v", err)
		}

		// masukkan kedalam slice bukus
		details = append(details, detail)

	}

	// return empty buku atau jika error
	return details, err
}

// // mengambil satu buku
// func GetOneDetail(id int64) ([]Detailed, error) {
// 	// mengkoneksikan ke db postgres
// 	db := config.CreateConnection()

// 	// kita tutup koneksinya di akhir proses
// 	defer db.Close()

// 	var details []Detailed

// 	// kita buat select query
// 	sqlStatement := `SELECT * FROM detailed WHERE id=$1`

// 	// mengeksekusi sql query
// 	rows, err := db.QueryRow(sqlStatement, id)

// 	if err != nil {
// 		log.Fatalf("Query could not be executed. %v", err)
// 	}

// 	// kita tutup eksekusi proses sql qeurynya
// 	defer rows.Close()

// 	// kita iterasi mengambil datanya
// 	for rows.Next() {
// 		var detail Detailed

// 		// kita ambil datanya dan unmarshal ke structnya
// 		err = rows.Scan(&detail.Id, &detail.Season, &detail.Episodes, &detail.Year)

// 		if err != nil {
// 			log.Fatalf("No data. %v", err)
// 		}

// 		// masukkan kedalam slice bukus
// 		details = append(details, detail)

// 	}

// 	// return empty buku atau jika error
// 	return details, err
// }

func DeleteDetailed(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM detailed WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Could not execute query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Data not found. %v", err)
	}

	fmt.Printf("Row(s) deleted %v", rowsAffected)

	return rowsAffected
}
