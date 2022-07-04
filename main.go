package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ClaudioGuevaraDev/Go-CLI-Cobra-MySQL/commands"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "contrasena",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "goclidb",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		os.Exit(1)
	}
	fmt.Println("Database connected")

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "CRUD with Cobra and MySQL",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.AddCommand(commands.InsertTaskCommand)
	rootCmd.AddCommand(commands.ListTasksCommand)
	rootCmd.AddCommand(commands.GetTaskCommand)
	rootCmd.AddCommand(commands.UpdateTaskCommand)
	rootCmd.AddCommand(commands.DeleteTaskCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
