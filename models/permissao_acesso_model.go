package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type PermissaoAcessoModel struct {
	Id        uint
	AtivoId   uint
	UsuarioId uint
}

func (m *PermissaoAcessoModel) TableName() string {
	return "usuarios_ativos"
}

func (m PermissaoAcessoModel) ToEntity() *entity.PermissaoAcesso {
	return &entity.PermissaoAcesso{
		Ativo:   entity.NewAtivo(m.AtivoId),
		Usuario: entity.NewUsuario(m.UsuarioId),
	}
}
