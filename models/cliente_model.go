package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type ClienteModel struct {
	Id        uint
	Nome      string
	Descricao string
	Status    string
}

func (m *ClienteModel) TableName() string {
	return "clientes"
}

func (m *ClienteModel) ToEntity() *entity.Cliente {
	return entity.
		NewClienteBuilder(m.Id).
		WithNome(m.Nome).
		WithStatus(entity.ClienteStatus(m.Status)).
		WithDescricao(m.Descricao).
		Build()
}
