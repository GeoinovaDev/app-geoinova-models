package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type LicencaTipoModel struct {
	Id        uint
	Nome      string
	Descricao string
}

func (m *LicencaTipoModel) TableName() string {
	return "licencas_tipos"
}

func (m *LicencaTipoModel) ToEntity() *entity.LicencaTipo {
	return entity.
		NewLicencaTipoBuilder(m.Id).
		WithNome(m.Nome).
		WithDescricao(m.Descricao).
		Build()
}
