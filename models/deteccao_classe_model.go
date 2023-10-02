package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type DeteccaoClasseModel struct {
	Id    uint
	Nome  string
	Tipo  string
	Color string
}

func (m DeteccaoClasseModel) TableName() string {
	return "deteccoes_classes"
}

func (m DeteccaoClasseModel) ToEntity() *entity.DeteccaoClasse {
	return entity.
		NewDeteccaoClasseBuilder(m.Id).
		WithNome(m.Nome).
		WithTipo(m.Tipo).
		WithColor(m.Color).
		Build()
}
