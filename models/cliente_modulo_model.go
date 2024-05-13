package models

type ClienteModuloModel struct {
	Id        uint
	Modulo    *ModuloModel  `gorm:"foreignKey:modulo_id"`
	ModuloId  uint          `gorm:"column:modulo_id"`
	Cliente   *ClienteModel `gorm:"foreignKey:cliente_id"`
	ClienteId uint          `gorm:"column:cliente_id"`
}

func (m *ClienteModuloModel) TableName() string {
	return "clientes_modulos"
}
