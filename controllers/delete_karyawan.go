package controllers

import (
    "database/sql"
    "net/http"
)

// Perbaiki nama fungsi dari NewUpadateKaryawan menjadi NewUpdateKaryawan
func NewDeleteKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM employee WHERE id = ?",id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
        
        http.Redirect(w,r,"/karyawan",http.StatusMovedPermanently)
		
		
	}
}