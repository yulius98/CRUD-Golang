package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewCreateKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            r.ParseForm()

            nama := r.Form["nama"][0]
            npwp := r.Form["npwp"][0]
            alamat := r.Form["alamat"][0]
            _, err := db.Exec("INSERT INTO employee (nama, npwp, alamat) VALUES (?,?,?)",nama, npwp, alamat)
            if err != nil {
                w.Write([]byte(err.Error()))
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            // Fix: Use proper redirect status code
            http.Redirect(w,r,"/karyawan",http.StatusSeeOther)

            return
        case "GET":
            fp := filepath.Join("views", "create.html")
            temp, err := template.ParseFiles(fp)
            if err != nil {
                w.Write([]byte(err.Error()))
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            
            err = temp.Execute(w, nil)
            if err != nil {
                
                w.Write([]byte(err.Error()))
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
        default:
            // Handle other HTTP methods if needed
            
            w.WriteHeader(http.StatusMethodNotAllowed)
            
        }

        
    }
}