package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ClaudioGuevaraDev/Go-CLI-Cobra-MySQL/database"
	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

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
		var tasks []Task

		sql := "SELECT * FROM tasks"
		res, err := database.DB.Query(sql)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for res.Next() {
			var task Task
			if err := res.Scan(&task.ID, &task.Title, &task.Description); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tasks = append(tasks, task)
		}

		table := simpletable.New()

		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "Title"},
				{Align: simpletable.AlignCenter, Text: "Description"},
			},
		}

		for _, row := range tasks {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(row.ID)},
				{Align: simpletable.AlignLeft, Text: row.Title},
				{Align: simpletable.AlignLeft, Text: row.Description},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}

		var StyleDefault = &simpletable.Style{
			Border: &simpletable.BorderStyle{
				TopLeft:            "+",
				Top:                "-",
				TopRight:           "+",
				Right:              "|",
				BottomRight:        "+",
				Bottom:             "-",
				BottomLeft:         "+",
				Left:               "|",
				TopIntersection:    "+",
				BottomIntersection: "+",
			},
			Divider: &simpletable.DividerStyle{
				Left:         "+",
				Center:       "-",
				Right:        "+",
				Intersection: "+",
			},
			Cell: "|",
		}
		table.SetStyle(StyleDefault)

		fmt.Println(table.String())
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
