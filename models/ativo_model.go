package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type AtivoModel struct {
	Id          uint
	Nome        string
	Wkt         string
	Status      string
	ClienteId   uint                 `gorm:"column:cliente_id"`
	Cliente     *ClienteModel        `gorm:"foreignKey:cliente_id"`
	CategoriaId *uint                `gorm:"column:categoria_id"`
	Categoria   *AtivoCategoriaModel `gorm:"foreignKey:categoria_id"`
}

func (m *AtivoModel) TableName() string {
	return "ativos"
}

func (m *AtivoModel) ToEntity() *entity.Ativo {
	cliente := entity.NewCliente(m.ClienteId)
	if m.Cliente != nil {
		cliente = m.Cliente.ToEntity()
	}

	var categoria *entity.AtivoCategoria
	if m.CategoriaId != nil {
		categoria = entity.NewAtivoCategoria(*m.CategoriaId, "")
	}

	if m.Categoria != nil {
		categoria = m.Categoria.ToEntity()
	}

	return entity.
		NewAtivoBuilder(m.Id).
		WithNome(m.Nome).
		WithWkt(m.Wkt).
		WithStatus(entity.AtivoStatus(m.Status)).
		WithCliente(cliente).
		WithCategoria(categoria).
		Build()
}
