package main // Menandakan ini program utama (entry point)

import (
	"net/http"

	"github.com/RESKY753/CRUD_GO/database" // Import package database buatan sendiri
	"github.com/RESKY753/CRUD_GO/routes"
)

func main() { // Fungsi utama yang pertama kali dijalankan
	db := database.InitDatabase() // Panggil fungsi InitDatabase dan simpan koneksi DB ke variabel db
	defer db.Close() // Tutup koneksi database otomatis saat program selesai

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}