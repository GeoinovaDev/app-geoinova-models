package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type PermissaoCategoriaModel struct {
	Id   uint
	Nome string
}

func (m *PermissaoCategoriaModel) TableName() string {
	return "permissoes_categorias"
}

func (m PermissaoCategoriaModel) ToEntity() *entity.PermissaoCategoria {
	return entity.
		NewPermissaoCategoriaBuilder(m.Id).
		WithNome(m.Nome).
		Build()
}
