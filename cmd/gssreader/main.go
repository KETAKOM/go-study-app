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
		r.LogicalName = table.GetStringValue(i, 1)
		r.PhysicalName = table.GetStringValue(i, 2)
		r.Type = table.GetStringValue(i, 3)
		r.NullAble = table.GetStringValue(i, 4)
		r.PK = table.GetStringValue(i, 5)
		r.Default = table.GetStringValue(i, 6)

		ts = append(ts, &r)
	}

	var sql string

	sql = "create table aaaa ("

	i := len(ts) - 1
	for s, t := range ts {
		if s != 0 {
			sql += t.LogicalName + " "
			sql += t.Type

			if s < i {
				sql += ", "
			}
		}
	}
	sql += ");"

	fmt.Println(sql)
}

type Schema struct {
	ID           string
	LogicalName  string
	PhysicalName string
	Type         string
	NullAble     string
	PK           string
	Default      string
}
