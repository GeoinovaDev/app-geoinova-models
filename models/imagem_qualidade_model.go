package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type ImagemQualidade struct {
	Id         uint
	Nome       string
	Color      string
	Background string
	Tipo       uint
}

func (m ImagemQualidade) TableName() string {
	return "imagens_qualidades"
}

func (m ImagemQualidade) ToEntity() *entity.ImagemQualidade {
	return entity.
		NewImagemQualidadeBuilder(m.Id).
		WithNome(m.Nome).
		WithColor(m.Color).
		WithBackground(m.Background).
		WithTipo(m.Tipo).
		Build()
}
