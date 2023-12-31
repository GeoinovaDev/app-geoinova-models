package models

import (
	"github.com/GeoinovaDev/app-geoinova-encrypt/encrypt"
	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type UsuarioModel struct {
	Id        uint
	Nome      string
	Email     string
	Senha     string
	Status    string
	ClienteId uint          `gorm:"column:cliente_id"`
	Cliente   *ClienteModel `gorm:"foreignKey:cliente_id"`
}

func (m *UsuarioModel) TableName() string {
	return "usuarios"
}

func (m *UsuarioModel) ToEntity() *entity.Usuario {
	cliente := entity.NewCliente(m.ClienteId)
	if m.Cliente != nil {
		cliente = m.Cliente.ToEntity()
	}

	return entity.
		NewUsuarioBuilder(m.Id).
		WithNome(m.Nome).
		WithEmail(m.Email).
		WithSenha(encrypt.DecodeString(m.Senha)).
		WithStatus(entity.UsuarioStatus(m.Status)).
		WithCliente(cliente).
		Build()
}
