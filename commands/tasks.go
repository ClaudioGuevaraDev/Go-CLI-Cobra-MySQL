package commands

import (
	"fmt"
	"os"

	"github.com/ClaudioGuevaraDev/Go-CLI-Cobra-MySQL/database"
	"github.com/ClaudioGuevaraDev/Go-CLI-Cobra-MySQL/helpers"
	"github.com/spf13/cobra"
)

var InsertTaskCommand = &cobra.Command{
	Use:   "insert",
	Short: "Insert a task",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		description := args[1]

		sql := "INSERT INTO tasks (title, description) VALUES (?, ?)"
		_, err := database.DB.Exec(sql, title, description)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Task created")
	},
}

var ListTasksCommand = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []helpers.Task

		sql := "SELECT * FROM tasks"
		res, err := database.DB.Query(sql)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for res.Next() {
			var task helpers.Task
			if err := res.Scan(&task.ID, &task.Title, &task.Description); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tasks = append(tasks, task)
		}

		helpers.CreateTable(tasks)
	},
}

var GetTaskCommand = &cobra.Command{
	Use:   "get",
	Short: "Get a task",
}

var UpdateTaskCommand = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
}

var DeleteTaskCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
}
