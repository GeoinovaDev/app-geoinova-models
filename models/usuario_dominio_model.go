package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type UsuarioDominioModel struct {
	Id         uint
	Dominio    string
	Role       string
	Checked    bool
	AppVersion string        `gorm:"column:app_version"`
	ApiVersion string        `gorm:"column:api_version"`
	UsuarioId  uint          `gorm:"column:usuario_id"`
	Usuario    *UsuarioModel `gorm:"foreignKey:usuario_id"`
}

func (m *UsuarioDominioModel) TableName() string {
	return "usuarios_dominios"
}

func (m *UsuarioDominioModel) ToEntity() *entity.UsuarioDominio {
	usuario := entity.NewUsuario(m.UsuarioId)
	if m.Usuario != nil {
		usuario = m.Usuario.ToEntity()
	}

	return entity.
		NewUsuarioDominioBuilder(m.Id).
		WithDominio(m.Dominio).
		WithRole(m.Role).
		WithChecked(m.Checked).
		WithUsuario(usuario).
		WithAppVersion(m.AppVersion).
		WithApiVersion(m.ApiVersion).
		Build()
}
