package app

import "database/sql"

func NewDB() (*sql.DB, error) {
	dbDriver := "mysql" // dbDriver merupakan tipe driver database yang digunakan (dalam kasus ini adalah MySQL).
	dbUser := "eep"     // dbUser dan dbPass adalah username dan password yang digunakan untuk masuk ke database.
	dbPass := "1903"
	dbName := "RESTfulAPI_todos" // dbName adalah nama database yang akan diakses.

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName) // sql.Open digunakan untuk membuka koneksi ke database dengan menggunakan driver, nama pengguna, password, dan nama database yang telah diinisialisasi sebelumnya.
	// Jika terjadi error selama proses pembukaan koneksi, maka akan terjadi panik (panic) dan program akan berhenti dengan mengeluarkan pesan error.
	if err != nil {
		return nil, err
	}

	return db, nil
}
