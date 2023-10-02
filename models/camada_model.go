package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type CamadaModel struct {
	Id          uint
	Nome        string
	Wkt         string
	Detalhe     string
	Area        float32
	AtivoId     uint        `gorm:"column:ativo_id"`
	Ativo       *AtivoModel `gorm:"foreignKey:ativo_id"`
	CategoriaId uint        `gorm:"column:categoria_id"`
}

func (m *CamadaModel) TableName() string {
	return "camadas"
}

func (m *CamadaModel) ToEntity() *entity.Camada {
	ativo := entity.NewAtivo(m.AtivoId)
	if m.Ativo != nil {
		ativo = m.Ativo.ToEntity()
	}

	categoria := entity.NewCamadaCategoria(m.CategoriaId)
	return entity.
		NewCamadaBuilder(m.Id).
		WithNome(m.Nome).
		WithWkt(m.Wkt).
		WithDetalhe(m.Detalhe).
		WithArea(m.Area).
		WithAtivo(ativo).
		WithCategoria(categoria).
		Build()
}
