package models

type PermissaoUsuarioModel struct {
	Id          uint
	Usuario     *UsuarioModel   `gorm:"foreignKey:usuario_id"`
	UsuarioId   uint            `gorm:"column:usuario_id"`
	Permissao   *PermissaoModel `gorm:"foreignKey:permissao_id"`
	PermissaoId uint            `gorm:"column:permissao_id"`
}

func (m *PermissaoUsuarioModel) TableName() string {
	return "permissoes_usuarios"
}
