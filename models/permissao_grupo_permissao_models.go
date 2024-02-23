package models

type PermissaoGrupoPermissaoModel struct {
	Id          uint
	GrupoId     uint `gorm:"column:grupo_id"`
	PermissaoId uint `gorm:"column:permissao_id"`
}

func (m *PermissaoGrupoPermissaoModel) TableName() string {
	return "permissoes_grupos_permissoes"
}
