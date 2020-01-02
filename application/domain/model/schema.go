package model

import (
	"errors"

	"github.com/KETAKOM/go-study-app/application/domain/model/schema"
	"github.com/yokoe/herschel"
)

type Table struct {
	Name string
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

func (*Schema) GetSchema(table *herschel.Table) ([]*Schema, error) {
	var ts []*Schema
	rows := table.GetRows()
	for i := 2; i < rows; i++ {
		var r Schema
		r.ID = table.GetStringValue(i, 0)
		r.LogicalName = table.GetStringValue(i, 1)
		r.PhysicalName = table.GetStringValue(i, 2)
		r.Type = table.GetStringValue(i, 3)
		r.NullAble = table.GetStringValue(i, 4)
		r.PK = table.GetStringValue(i, 5)
		r.Default = table.GetStringValue(i, 6)

		if !schema.ValidateType(r.Type) {
			return []*Schema{}, errors.New("validate failed")
		}

		ts = append(ts, &r)
	}
	return ts, nil
}
