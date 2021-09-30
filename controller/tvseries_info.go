package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv" // package yang digunakan untuk mengubah string menjadi tipe int

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Buku didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []models.TV `json:"data"`
}

// Add new TV series
func NewTV(w http.ResponseWriter, r *http.Request) {

	// create an empty user of type models.User
	// kita buat empty buku dengan tipe models.Buku
	var tv models.TV

	// decode data json request ke buku
	err := json.NewDecoder(r.Body).Decode(&tv)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert buku
	insertID := models.AddTV(tv)

	// format response objectnya
	res := response{
		ID:      insertID,
		Message: "TV series info has been added",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// return one row of tv series by id
func GetTV(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan idbuku dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("String to int conversion error.  %v", err)
	}

	// memanggil models ambilsatubuku dengan parameter id yg nantinya akan mengambil single data
	tv, err := models.GetOneTV(int64(id))

	if err != nil {
		log.Fatalf("TV series info could not be restracted. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(tv)
}

// return all rows from tvseries_info
func GetTVAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaBuku
	tvs, err := models.GetTVdata()

	if err != nil {
		log.Fatalf("Data could not be restracted. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = tvs

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

//Update one row from tvseries_info by id
func UpdateTVNew(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("TV series info could not be restracted.  %v", err)
	}

	// buat variable buku dengan type models.Buku
	var tv models.TV

	// decode json request ke variable buku
	err = json.NewDecoder(r.Body).Decode(&tv)

	if err != nil {
		log.Fatalf("Could not decode body request.  %v", err)
	}

	// panggil updatebuku untuk mengupdate data
	updatedRows := models.UpdateTV(int64(id), tv)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("TV series updated. Updated %v row(s)/record(s)", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

//Delete one row
func RemoveTV(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error during string to int conversion.  %v", err)
	}

	// panggil fungsi hapusbuku , dan convert int ke int64
	deletedRows := models.DeleteTV(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Tv series info deleted. Affected Row(s) :  %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
