package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type LicencaTipoModel struct {
	Id   uint
	Nome string
}

func (l *LicencaTipoModel) TableName() string {
	return "licencas_tipos"
}

func (l *LicencaTipoModel) ToEntity() *entity.LicencaTipo {
	return entity.NewLicencaTipoWithNome(l.Id, l.Nome)
}
