package main

import (
	"fmt"
	"os"

	"github.com/yokoe/herschel"
	"github.com/yokoe/herschel/option"
)

func main() {
	dir := os.Getenv("JSON_DIR")
	op := option.WithServiceAccountCredentials(dir)
	client, err := herschel.NewClient(op)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := os.Getenv("SHEET_ID")
	name := os.Getenv("SHEET_NAME")

	table, err := client.ReadTable(id, name)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := table.GetRows()

	var ts []*Schema
	for i := 1; i < rows; i++ {
		var r Schema
		r.ID = table.GetStringValue(i, 0)
		r.Field = table.GetStringValue(i, 1)
		r.Type = table.GetStringValue(i, 2)

		ts = append(ts, &r)
	}

	for _, t := range ts {
		fmt.Println(t)
	}
}

type Schema struct {
	ID    string
	Field string
	Type  string
}
