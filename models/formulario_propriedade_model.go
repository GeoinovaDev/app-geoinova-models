package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type FormularioPropriedadeModel struct {
	Id   uint
	Nome string
	Tipo string
}

func (m *FormularioPropriedadeModel) TableName() string {
	return "formularios_propriedades"
}

func (m *FormularioPropriedadeModel) ToEntity() *entity.FormularioPropriedade {
	return entity.
		NewFormularioPropriedadeBuilder(m.Id).
		WithNome(m.Nome).
		WithTipo(entity.FormularioPropriedadeTipo(m.Tipo)).
		Build()
}
