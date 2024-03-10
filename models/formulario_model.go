package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type FormularioModel struct {
	Id        uint
	Nome      string
	ClienteId uint          `gorm:"column:cliente_id"`
	Cliente   *ClienteModel `gorm:"foreignKey:cliente_id"`
	IsPadrao  bool          `gorm:"column:is_default"`
}

func (m *FormularioModel) TableName() string {
	return "formularios"
}

func (m *FormularioModel) ToEntity() *entity.Formulario {
	cliente := entity.NewCliente(m.ClienteId)
	if m.Cliente != nil {
		cliente = m.Cliente.ToEntity()
	}

	return entity.
		NewFormularioBuilder(m.Id).
		WithNome(m.Nome).
		WithIsPadrao(m.IsPadrao).
		WithCliente(cliente).
		Build()
}
