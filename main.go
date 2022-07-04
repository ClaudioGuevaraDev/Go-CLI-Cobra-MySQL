package main

import (
	"fmt"
	"os"

	"github.com/ClaudioGuevaraDev/Go-CLI-Cobra-MySQL/commands"
	"github.com/spf13/cobra"
)

func main() {
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
