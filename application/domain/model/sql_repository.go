package model

type SQLRepository interface {
	GetSQL(ts []*Schema) *SQL
}
