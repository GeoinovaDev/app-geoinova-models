package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type FormularioPropriedadeModel struct {
	Id           uint
	Nome         string
	Field        string
	Tipo         string
	FormularioId uint `gorm:"column:formulario_id"`
}

func (m *FormularioPropriedadeModel) TableName() string {
	return "formularios_propriedades"
}

func (m *FormularioPropriedadeModel) ToEntity() *entity.FormularioPropriedade {
	return entity.
		NewFormularioPropriedadeBuilder(m.Id).
		WithNome(m.Nome).
		WithField(m.Field).
		WithTipo(entity.FormularioPropriedadeTipo(m.Tipo)).
		Build()
}
