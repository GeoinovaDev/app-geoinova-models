package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type ModuloModel struct {
	Id        uint
	Nome      string
	Descricao string
}

func (m ModuloModel) TableName() string {
	return "modulos"
}

func (m ModuloModel) ToEntity() *entity.Modulo {
	return entity.
		NewModuloBuilder(m.Id).
		WithNome(m.Nome).
		Build()
}
