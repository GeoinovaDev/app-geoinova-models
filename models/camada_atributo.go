package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type CamadaAtributoModel struct {
	Id       uint
	Nome     string
	Valor    string
	CamadaId uint `gorm:"column:camada_id"`
}

func (m *CamadaAtributoModel) TableName() string {
	return "camadas_atributos"
}

func (m *CamadaAtributoModel) ToEntity() *entity.CamadaAtributo {
	return entity.
		NewCamadaAtributoBuilder(m.Id).
		WithNome(m.Nome).
		WithValor(m.Valor).
		WithCamada(entity.NewCamada(m.CamadaId)).
		Build()
}
