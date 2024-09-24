package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type IncendioConfigModel struct {
	Id        uint
	ClienteId *uint
	Nome      string
	Raio      float64
}

func (m *IncendioConfigModel) TableName() string {
	return "incendios_configuracoes"
}

func (m *IncendioConfigModel) ToEntity() *entity.IncendioConfig {
	cliente := entity.NewCliente(*m.ClienteId)
	if m.ClienteId != nil {
		cliente = entity.NewCliente(*m.ClienteId)
	}

	return entity.NewIncendioConfigBuilder(m.Id).
		WithNome(m.Nome).
		WithRaio(m.Raio).
		WithCliente(cliente).
		Build()
}
