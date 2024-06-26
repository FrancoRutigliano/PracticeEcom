package main

import (
	"database/sql"
	"log"

	"github.com/FrancoRutigliano/cmd/api"
	"github.com/FrancoRutigliano/internal/config"
	"github.com/FrancoRutigliano/internal/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	PORT := ":8080"
	server := api.NewAPIServer(PORT, db)
	if err := server.Run(); err != nil {
		log.Fatal("Error to initialize the server on port: ", PORT)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Succesfully conected")
}
