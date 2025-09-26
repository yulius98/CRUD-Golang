package controllers

import (
    "database/sql"
    "html/template"
    "net/http"
    "path/filepath"
)

// Perbaiki nama fungsi dari NewUpadateKaryawan menjadi NewUpdateKaryawan
func NewUpdateKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            r.ParseForm()

            nama := r.Form["nama"][0]
            npwp := r.Form["npwp"][0]
            alamat := r.Form["alamat"][0]
            // Perbaiki SQL query - UPDATE tidak menggunakan INTO
            _, err := db.Exec("UPDATE employee SET nama=?, npwp=?, alamat=? WHERE id=?", nama, npwp, alamat, r.Form["id"][0])
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
                return
            }
            http.Redirect(w,r,"/karyawan",http.StatusSeeOther)
            return

        case "GET":
            id := r.URL.Query().Get("id")
            row := db.QueryRow("SELECT nama, npwp, alamat FROM employee WHERE id =?",id)

            var karyawan Karyawan
            err := row.Scan(
                &karyawan.Nama,
                &karyawan.NPWP,
                &karyawan.Alamat,
            )
            karyawan.Id = id
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
                return
            }

            fp := filepath.Join("views", "update.html")
            temp, err := template.ParseFiles(fp)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
                return
            }
            
            data := make(map[string]any)
            data["karyawan"] = karyawan
            
            // Perbaiki - gunakan data sebagai parameter, bukan nil
            err = temp.Execute(w, data)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
                return
            }
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    }
}