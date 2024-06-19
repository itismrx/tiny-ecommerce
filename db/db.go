package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN()) //"root:root@tcp(mysql:3306)/GoLife"

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InitalStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db is connected succcessfully")
}
