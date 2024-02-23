package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type PermissaoModel struct {
	Id       uint
	Nome     string
	Leitura  bool
	Escrita  bool
	Exclusao bool
}

func (m *PermissaoModel) TableName() string {
	return "permissoes"
}

func (m PermissaoModel) ToEntity() *entity.Permissao {
	return entity.
		NewPermissaoBuilder(m.Id).
		WithNome(m.Nome).
		WithLeitura(m.Leitura).
		WithEscrita(m.Escrita).
		WithExclusao(m.Exclusao).
		Build()
}
