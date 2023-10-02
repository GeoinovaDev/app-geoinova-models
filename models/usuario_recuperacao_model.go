package models

import (
	"time"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type UsuarioRecuperacaoModel struct {
	Id        uint
	Hash      string
	UsuarioId uint          `gorm:"column:usuario_id"`
	Usuario   *UsuarioModel `gorm:"foreignKey:usuario_id"`
	CreatedAt []uint8       `gorm:"column:created_at"`
	UpdatedAt []uint8       `gorm:"column:updated_at"`
}

func (m *UsuarioRecuperacaoModel) TableName() string {
	return "usuarios_recuperacoes"
}

func (m *UsuarioRecuperacaoModel) ToEntity() *entity.UsuarioRecuperacao {
	usuario := entity.NewUsuario(m.UsuarioId)
	if m.Usuario != nil {
		usuario = m.Usuario.ToEntity()
	}

	createdAt, _ := time.Parse("2006-01-02 15:04:05", string(m.CreatedAt))
	return entity.NewUsuarioRecuperacao(m.Id, usuario, m.Hash, createdAt)
}
