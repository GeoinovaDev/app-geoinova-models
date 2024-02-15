package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type FormularioModel struct {
	Id   uint
	Nome string
}

func (m *FormularioModel) TableName() string {
	return "formularios"
}

func (m *FormularioModel) ToEntity() *entity.Formulario {
	return entity.
		NewFormularioBuilder(m.Id).
		WithNome(m.Nome).
		Build()
}
