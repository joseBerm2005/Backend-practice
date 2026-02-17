package main

// API's entry point

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/joseBerm2005/Backend-practice/cmd/api"
	"github.com/joseBerm2005/Backend-practice/config"
	"github.com/joseBerm2005/Backend-practice/db"
)

// creates server and runs the Run function
func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8000", db)

	if err := server.Run(); err != nil {
		log.Fatal()
	}
}

func initStorage( db *sql.DB) {
	err := db.Ping()
	
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: successfully connected")
}