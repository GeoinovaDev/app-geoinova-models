package models

type PermissaoGrupoPermissaoModel struct {
	Id               uint
	GrupoPermissaoId uint `gorm:"column:grupo_permissao_id"`
	PermissaoId      uint `gorm:"column:permissao_id"`
	Enabled          bool
}

func (m *PermissaoGrupoPermissaoModel) TableName() string {
	return "permissoes_grupos_permissoes"
}
