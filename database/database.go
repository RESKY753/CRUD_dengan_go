package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB { //sql.db itu poiternya
	koneksi := "root@tcp(localhost:3306)/crud_go" //koneksi dalam go

	db, err := sql.Open("mysql", koneksi) 
	if err != nil { //kondisi untuk mengecek
		panic(err)
	}

	err = db.Ping() //fungsi ping ini untuk mengembalikan error
	if err != nil {
		panic(err)
	}
	return db
}