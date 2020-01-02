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

func (uc *CreateSQLUseCase) CreateSQL(table *herschel.Table) (*model.SQL, error) {
	schama, err := uc.SchemaRepository.GetSchema(table)
	if err != nil {
		return nil, err
	}

	sql := uc.SQLRepository.GetSQL(schama)

	return sql, nil
}
