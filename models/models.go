package models

import (
	"database/sql"
	"fmt"
	"go-postgresql/config"
	"log"

	_ "github.com/lib/pq"
)

//buku schema dari tabel buku
//kita coba dengan jika datanya null
//jika return datanya ada yang null, silahkan pake nullstring, contohnya di bawah
// penulis config.Nullstring `json:"penulis"`
type Buku struct {
	ID            int64  `json:"id"`
	Judul_buku    string `json:"judul_buku"`
	Penulis       string `json:"penulis"`
	Tgl_publikasi string `json:"tgl_publikasi"`
}

func TambahBuku(buku Buku) int64 {
	//mengkoneksikan ke db postgresql
	db := config.CreateConnection()
	//kita tutup koneksinya di akhir proses
	defer db.Close()

	//kita buat insert query
	//mengembalikan nilai id akan mengembalikan id dari buku yang dimasukkan ke db
	sqlStatement := `INSERT INTO buku (judul_buku, penulis, tgl_publikasi) VALUES ($1, $2, $3) RETURNING id`

	//id yang dimasukkan akan disimpan di id ini
	var id int64

	//scan function akan menyimpan insert id di dalam id id
	err := db.QueryRow(sqlStatement, buku.Judul_buku, buku.Penulis, buku.Tgl_publikasi).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	//return insert id
	return id
}

//ambil satu buku
func AmbilSemuaBuku() ([]Buku, error) {
	//mengkoneksikan ke db postgres
	db := config.CreateConnection()
	//kita tutup koneksinya di akhir proses
	defer db.Close()

	var bukus []Buku
	//kita buat select query
	sqlStatement := `SELECT * FROM buku`

	//mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	//kita tutup eksekusi proses sql querynya
	defer rows.Close()

	//kite iterasi mengambil datanya
	for rows.Next() {
		var buku Buku
		//kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&buku.ID, &buku.Judul_buku, &buku.Penulis, &buku.Tgl_publikasi)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		//masukkan kedalam slice bukus
		bukus = append(bukus, buku)

	}

	//return empty buku atau jika error
	return bukus, err
}

//mengambil satu buku
func AmbilSatuBuku(id int64) (Buku, error) {
	//mengkoneksikan  ke db postgres
	db := config.CreateConnection()

	//kita tutup koneksinya di akhir proses
	defer db.Close()

	var buku Buku

	//buat sql query
	sqlStatement := `SELECT * FROM buku WHERE id=$1`

	//eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&buku.ID, &buku.Judul_buku, &buku.Penulis, &buku.Tgl_publikasi)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return buku, nil
	case nil:
		return buku, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return buku, err
}

// update user in the DB
func UpdateBuku(id int64, buku Buku) int64 {
	//mengkoneksikan ke db postgres
	db := config.CreateConnection()
	//kita tutup koneksinya di akhir proses
	defer db.Close()
	//kita buat sql query create
	sqlStatement := `UPDATE buku SET judul_buku=$2, penulis=$3, tgl_publikasi=$4 WHERE id=$1`

	//eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, buku.Judul_buku, buku.Penulis, buku.Tgl_publikasi)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	//cek berapa banyak row/data yang di update
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengechek rows/data yang di update. %v", err)

	}
	fmt.Printf("Total rows/record yang di update %v\n", rowsAffected)

	return rowsAffected
}

func HapusBuku(id int64) int64 {
	//mengkoneksikan ke db postgres
	db := config.CreateConnection()

	//kita tutup koneksinya di akhir proses
	defer db.Close()

	//buat sql query
	sqlStatement := `DELETE FROM buku WHERE id=$1`

	//ekseskusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	//cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)

	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
