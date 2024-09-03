package models

import (
	"github.com/GeoinovaDev/app-geoinova-encrypt/encrypt"
	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type UsuarioModel struct {
	Id               uint
	Nome             string
	Email            string
	Senha            string
	Status           string
	Descricao        string
	Papel            string
	Telefone         *string
	Notificacao      string
	ClienteId        uint                 `gorm:"column:cliente_id"`
	Cliente          *ClienteModel        `gorm:"foreignKey:cliente_id"`
	GrupoPermissaoId uint                 `gorm:"column:grupo_permissao_id"`
	GrupoPermissao   *GrupoPermissaoModel `gorm:"foreignKey:grupo_permissao_id"`
}

func (m *UsuarioModel) TableName() string {
	return "usuarios"
}

func (m *UsuarioModel) ToEntity() *entity.Usuario {
	cliente := entity.NewCliente(m.ClienteId)
	if m.Cliente != nil {
		cliente = m.Cliente.ToEntity()
	}
	grupoPermissao := entity.NewGrupoPermissao(m.GrupoPermissaoId)
	if m.GrupoPermissao != nil {
		grupoPermissao = m.GrupoPermissao.ToEntity()
	}

	telefone := ""
	if m.Telefone != nil {
		telefone = *m.Telefone
	}

	return entity.
		NewUsuarioBuilder(m.Id).
		WithNome(m.Nome).
		WithEmail(m.Email).
		WithPapel(entity.UsuarioPapel(m.Papel)).
		WithDescricao(m.Descricao).
		WithSenha(encrypt.DecodeString(m.Senha)).
		WithStatus(entity.UsuarioStatus(m.Status)).
		WithCliente(cliente).
		WithGrupoPermissao(grupoPermissao).
		WithTelefone(entity.ParseTelefone(telefone)).
		WithRecebeNotificacao(m.Notificacao == "A").
		Build()
}
