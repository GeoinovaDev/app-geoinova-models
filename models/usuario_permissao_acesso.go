package models

type UsuarioPermissaoAcesso struct {
	Usuario   *UsuarioModel `gorm:"foreignKey:usuario_id"`
	UsuarioId uint          `gorm:"column:usuario_id"`
	Ativo     *AtivoModel   `gorm:"foreignKey:ativo_id"`
	AtivoId   uint          `gorm:"column:ativo_id"`
}

func (m *UsuarioPermissaoAcesso) TableName() string {
	return "usuarios_ativos"
}
