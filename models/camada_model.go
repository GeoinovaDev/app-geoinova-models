package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type CamadaModel struct {
	Id           uint
	Nome         string
	Wkt          string
	Detalhe      string
	AtivoId      uint             `gorm:"column:ativo_id"`
	Ativo        *AtivoModel      `gorm:"foreignKey:ativo_id"`
	CategoriaId  uint             `gorm:"column:categoria_id"`
	Formulario   *FormularioModel `gorm:"foreignKey:formulario_id"`
	FormularioId *uint            `gorm:"column:formulario_id"`
}

func (m *CamadaModel) TableName() string {
	return "camadas"
}

func (m *CamadaModel) ToEntity() *entity.Camada {
	ativo := entity.NewAtivo(m.AtivoId)
	if m.Ativo != nil {
		ativo = m.Ativo.ToEntity()
	}

	var formulario *entity.Formulario
	if m.FormularioId != nil {
		formulario = entity.NewFormulario(*m.FormularioId)
	}

	builder := entity.
		NewCamadaBuilder(m.Id).
		WithNome(m.Nome).
		WithWkt(m.Wkt).
		WithDetalhe(m.Detalhe).
		WithAtivo(ativo).
		WithFormulario(formulario).
		WithCategoria(entity.NewCamadaCategoria(m.CategoriaId))

	return builder.Build()
}
