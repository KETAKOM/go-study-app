package repository

import (
	"github.com/KETAKOM/go-study-app/application/domain/model"
	"github.com/yokoe/herschel"
)

type SchemaRepository interface {
	GetSchema(table herschel.Table) []*model.Schema
}
