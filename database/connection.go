package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitializeConnection() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "contrasena",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "goclidb",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error to connect mysql by", err)
		os.Exit(1)
	}

	DB = db
}
