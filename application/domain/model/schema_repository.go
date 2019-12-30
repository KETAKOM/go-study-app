package model

import (
	"github.com/yokoe/herschel"
)

type SchemaRepository interface {
	GetSchema(table *herschel.Table) []*Schema
}
