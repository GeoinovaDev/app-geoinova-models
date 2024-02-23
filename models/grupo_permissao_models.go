package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type GrupoPermissaoModel struct {
	Id   uint
	Nome string
}

func (m *GrupoPermissaoModel) TableName() string {
	return "grupos_permissoes"
}

func (m *GrupoPermissaoModel) ToEntity() *entity.GrupoPermissao {
	return entity.
		NewGrupoPermissaoBuilder(m.Id).
		WithNome(m.Nome).
		Build()
}
