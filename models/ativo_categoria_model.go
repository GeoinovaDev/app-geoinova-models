package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type AtivoCategoriaModel struct {
	Id   uint
	Nome string
}

func (m AtivoCategoriaModel) TableName() string {
	return "ativos_categorias"
}

func (m AtivoCategoriaModel) ToEntity() *entity.AtivoCategoria {
	return entity.NewAtivoCategoria(m.Id, m.Nome)
}
