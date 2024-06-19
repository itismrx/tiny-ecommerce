package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/itismrx/tiny-ecommerce/cmd/api"
	"github.com/itismrx/tiny-ecommerce/config"
	"github.com/itismrx/tiny-ecommerce/db"
)

func main() {
	ldb, err := db.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBUserPassword,
		Addr:   config.Envs.PublicHost,
		DBName: config.Envs.DBName,
		// Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	db.InitalStorage(ldb)
	server := api.NewAPIServer(":8080", ldb)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
