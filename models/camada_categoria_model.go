package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type CamadaCategoriaModel struct {
	Id          uint
	Nome        string
	Cor         string `gorm:"column:cor_preenchimento"`
	Borda       string `gorm:"column:cor_borda"`
	CategoriaId uint   `gorm:"column:categoria_id"`
}

func (m *CamadaCategoriaModel) TableName() string {
	return "camadas_categorias"
}

func (m *CamadaCategoriaModel) ToEntity() *entity.CamadaCategoria {
	categoria := entity.NewCamadaCategoria(m.CategoriaId)
	if categoria.Id == 0 {
		categoria = nil
	}

	return entity.
		NewCamadaCategoriaBuilder(m.Id).
		WithNome(m.Nome).
		WithCor(m.Cor).
		WithBorda(m.Borda).
		WithCategoria(categoria).
		Build()
}
