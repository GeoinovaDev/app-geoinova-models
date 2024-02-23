package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type PermissaoModel struct {
	Id          uint
	Nome        string
	Action      string
	Leitura     bool
	Escrita     bool
	Exclusao    bool
	CategoriaId uint                     `gorm:"column:categoria_id"`
	Categoria   *PermissaoCategoriaModel `gorm:"foreignKey:categoria_id"`
}

func (m *PermissaoModel) TableName() string {
	return "permissoes"
}

func (m PermissaoModel) ToEntity() *entity.Permissao {
	var categoria *entity.PermissaoCategoria
	if m.CategoriaId != 0 {
		categoria = entity.NewPermissaoCategoria(m.CategoriaId)
	}
	
	return entity.
		NewPermissaoBuilder(m.Id).
		WithNome(m.Nome).
		WithCategoria(categoria).
		WithAction(m.Action).
		WithLeitura(m.Leitura).
		WithEscrita(m.Escrita).
		WithExclusao(m.Exclusao).
		Build()
}
