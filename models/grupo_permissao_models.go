package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type GrupoPermissaoModel struct {
	Id        uint
	Nome      string
	ClienteId uint          `gorm:"column:cliente_id"`
	Cliente   *ClienteModel `gorm:"foreignKey:cliente_id"`
}

func (m *GrupoPermissaoModel) TableName() string {
	return "grupos_permissoes"
}

func (m *GrupoPermissaoModel) ToEntity() *entity.GrupoPermissao {
	cliente := entity.NewCliente(m.ClienteId)
	if m.Cliente != nil {
		cliente = m.Cliente.ToEntity()
	}

	return entity.
		NewGrupoPermissaoBuilder(m.Id).
		WithNome(m.Nome).
		WithCliente(cliente).
		Build()
}
