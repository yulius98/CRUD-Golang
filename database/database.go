package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
)

func InitDatabase() *sql.DB {
	connDb := "root@tcp(localhost:3306)/db_hrd_golang"
	db, err := sql.Open("mysql", connDb)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}