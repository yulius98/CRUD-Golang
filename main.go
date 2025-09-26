package main

import (
	"apphrd/database"
	"apphrd/routes"
	"net/http"
)

func main() {
	db := database.InitDatabase()
	
	server := http.NewServeMux()
	routes.MapRoute(server, db)
	http.ListenAndServe(":8080",server)
}
