package controller

import (
	"encoding/json" //package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"go-postgresql/models" // models package dimana buku di definisikan
	"log"
	"net/http" //digunakan untuk mengakses objek permintaan dan respons dari api
	"strconv"  //package yang digunakan untuk mengubah string menjadi tipe int

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    //postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Buku `json:"data"`
}

//TambahBuku
func TmbhBuku(w http.ResponseWriter, r *http.Request) {
	//create an empty user of type models.User
	//kita buat empty buku dengan tipe models.Buku
	var buku models.Buku

	//decode data json request ke buku
	err := json.NewDecoder(r.Body).Decode(&buku)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body. %v", err)
	}
	//panggil modelnya lalu insert buku
	insertID := models.TambahBuku(buku)

	//format response objectnya
	res := response{
		ID:      insertID,
		Message: "Data buku telah ditambahkan",
	}

	//kirim response
	json.NewEncoder(w).Encode(res)
}

//AmbilBuku mengambil single data dengan parameter id
func AmbilBuku(w http.ResponseWriter, r *http.Request) {
	//kita set headernya
	w.Header().Set("Context-Type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//dapatkan idbuku dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	//konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	//memanggil models ambilsatubuku dengan parameter id yang nantinya akan mengambil single data
	buku, err := models.AmbilSatuBuku(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data buku. %v", err)
	}

	//kirim response
	json.NewEncoder(w).Encode(buku)

}

//Ambil semua data buku
func AmbilSemuaBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//memanggil models AmbilSemuaBuku
	bukus, err := models.AmbilSemuaBuku()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data .%v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = bukus
	//kirim semua response
	json.NewEncoder(w).Encode(response)
}
func UpdateBuku(w http.ResponseWriter, r *http.Request) {
	//kita ambil request parameter idnya
	params := mux.Vars(r)

	//konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}
	//buat variabel dengan type models.buku
	var buku models.Buku

	//decode json request ke variable buku
	err = json.NewDecoder(r.Body).Decode(&buku)

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}
	//panggil updatebuku untuk mengupdate data
	updatedRows := models.UpdateBuku(int64(id), buku)

	//ini adalah format message berupa string
	msg := fmt.Sprintf("Buku telah berhasil di update. Jumlah yang di update %v rows/record", updatedRows)
	//ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func HapusBuku(w http.ResponseWriter, r *http.Request) {
	//kita ambil request parameter idnya
	params := mux.Vars(r)

	//konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("TIdak bisa mengubah dari string ke int. %v", err)
	}

	//panggil fungsi hapusbuku, dan convert int ke int64
	deletedRows := models.HapusBuku(int64(id))

	//ini adalah format message berupa string
	msg := fmt.Sprintf("buku sukses di hapus. Total data yang di hapus %v", deletedRows)

	//ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	//send the response
	json.NewEncoder(w).Encode(res)
}
