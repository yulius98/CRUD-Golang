package routes

import (
	"apphrd/controllers"
	"database/sql"
	"net/http"
)

func MapRoute(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controllers.ConnServer() )
	server.HandleFunc("/karyawan", controllers.NewIndexKaryawan(db))
	server.HandleFunc("/karyawan/create", controllers.NewCreateKaryawan(db))
	server.HandleFunc("/karyawan/update", controllers.NewUpdateKaryawan(db))
	server.HandleFunc("/karyawan/delete", controllers.NewDeleteKaryawan(db))
}