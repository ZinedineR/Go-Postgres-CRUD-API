package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Buku schema dari tabel Buku
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type TV struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Producer string `json:"producer"`
}

func AddTV(tv TV) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari buku yang dimasukkan ke db
	sqlStatement := `INSERT INTO tvseries_info (title, producer) VALUES ($1, $2) RETURNING id`

	// id yang dimasukkan akan disimpan di id ini
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, tv.Title, tv.Producer).Scan(&id)

	if err != nil {
		log.Fatalf("Query could not be executed. %v", err)
	}

	fmt.Printf("Data inserted in record %v", id)

	// return insert id
	return id
}

// ambil buku
func GetTVdata() ([]TV, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var tvs []TV

	// kita buat select query
	sqlStatement := `SELECT * FROM tvseries_info`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Query could not be executed. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var tv TV

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&tv.Id, &tv.Title, &tv.Producer)

		if err != nil {
			log.Fatalf("No data. %v", err)
		}

		// masukkan kedalam slice bukus
		tvs = append(tvs, tv)

	}

	// return empty buku atau jika error
	return tvs, err
}

// mengambil satu buku
func GetOneTV(id int64) (TV, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var tv TV

	// buat sql query
	sqlStatement := `SELECT * FROM tvseries_info WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&tv.Id, &tv.Title, &tv.Producer)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Data not found")
		return tv, nil
	case nil:
		return tv, nil
	default:
		log.Fatalf("No data. %v", err)
	}

	return tv, err
}

// update user in the DB
func UpdateTV(id int64, tv TV) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE tvseries_info SET title=$2, producer=$3 WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, tv.Title, tv.Producer)

	if err != nil {
		log.Fatalf("Could not execute query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error checking the row. %v", err)
	}

	fmt.Printf("Row(s) affected %v\n", rowsAffected)

	return rowsAffected
}

func DeleteTV(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM tvseries_info WHERE id=$1`

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
