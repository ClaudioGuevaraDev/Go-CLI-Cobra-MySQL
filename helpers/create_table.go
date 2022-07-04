package helpers

import (
	"fmt"
	"strconv"

	"github.com/alexeyco/simpletable"
)

func CreateTable(tasks []Task) {
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
}
