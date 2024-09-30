package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type IncendioConfigModel struct {
	Id        uint
	ClienteId uint
	Nome      string
	Raio      float64
}

func (m *IncendioConfigModel) TableName() string {
	return "incendios_configuracoes"
}

func (m *IncendioConfigModel) ToEntity() *entity.IncendioConfig {
	return entity.NewIncendioConfigBuilder(m.Id).
		WithNome(m.Nome).
		WithRaio(m.Raio).
		WithCliente(entity.NewCliente(m.ClienteId)).
		Build()
}
