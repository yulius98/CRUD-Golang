package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Karyawan struct {
	Id string
	Nama string
	NPWP string
	Alamat string
}

func NewIndexKaryawan(db *sql.DB ) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, npwp, alamat FROM employee")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
		defer rows.Close()

		var karyawans []Karyawan
		for rows.Next() {
			var karyawan Karyawan

			err = rows.Scan(
				&karyawan.Id,
				&karyawan.Nama,
				&karyawan.NPWP,
				&karyawan.Alamat,

			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return	
			}


			karyawans = append(karyawans,karyawan)
		} 	

		fp := filepath.Join("views","index.html")
		temp, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}

		data := make(map[string]any)
		data["karyawans"] = karyawans

		err = temp.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
	}
}