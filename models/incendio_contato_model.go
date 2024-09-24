package models

type IncendioContato struct {
	Id               uint
	IncendioConfigId uint `gorm:"column:incendio_configuracao_id"`
	UsuarioId        uint `gorm:"column:usuario_id"`
	Email            string
	Telefone         string
}

func (m *IncendioContato) TableName() string {
	return "incendios_contatos"
}
