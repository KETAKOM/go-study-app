package usecase

import (
	"github.com/KETAKOM/go-study-app/application/domain/model"
	"github.com/yokoe/herschel"
)

type CreateSQLUseCase struct {
	SchemaRepository model.SchemaRepository
	SQLRepository    model.SQLRepository
}

func NewCreateSQLUseCase() *CreateSQLUseCase {
	return &CreateSQLUseCase{}
}

func (uc *CreateSQLUseCase) CreateSQL(table *herschel.Table) *model.SQL {
	schama := uc.SchemaRepository.GetSchema(table)
	return uc.SQLRepository.GetSQL(schama)
}
