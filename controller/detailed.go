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

type ResponseDetail struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.Detailed `json:"data"`
}

// Add Detailed
func NewDetailed(w http.ResponseWriter, r *http.Request) {

	// create an empty user of type models.User
	// kita buat empty buku dengan tipe models.Buku
	var detail models.Detailed

	// decode data json request ke buku
	err := json.NewDecoder(r.Body).Decode(&detail)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert buku
	insertID := models.AddDetailed(detail)

	// format response objectnya
	res := response{
		ID:      insertID,
		Message: "Detailed TV series info has been added",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// // AmbilBuku mengambil single data dengan parameter id
// func GetDetailedOne(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// memanggil models AmbilSemuaBuku
// 	details, err := models.GetOneDetail()

// 	if err != nil {
// 		log.Fatalf("Data could not be restracted. %v", err)
// 	}

// 	var response ResponseDetail
// 	response.Status = 1
// 	response.Message = "Success"
// 	response.Data = details

// 	// kirim semua response
// 	json.NewEncoder(w).Encode(response)
// }

// Return all detailed rows
func GetDetailedAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaBuku
	details, err := models.GetDetaildata()

	if err != nil {
		log.Fatalf("Data could not be restracted. %v", err)
	}

	var response ResponseDetail
	response.Status = 1
	response.Message = "Success"
	response.Data = details
	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

// Delete row from detailed
func RemoveDetailed(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error during string to int conversion.  %v", err)
	}

	// panggil fungsi hapusbuku , dan convert int ke int64
	deletedRows := models.DeleteDetailed(int64(id))

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
