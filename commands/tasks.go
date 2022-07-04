package commands

import "github.com/spf13/cobra"

var InsertTaskCommand = &cobra.Command{
	Use:   "insert",
	Short: "Insert a task",
}

var ListTasksCommand = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
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
