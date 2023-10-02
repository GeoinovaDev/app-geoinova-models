package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type AtivoMaisAcessadoModel struct {
	Id          uint
	Ativo       *AtivoModel   `gorm:"foreignKey:ativo_id"`
	AtivoId     uint          `gorm:"column:ativo_id"`
	Usuario     *UsuarioModel `gorm:"foreignKey:usuario_id"`
	UsuarioId   uint          `gorm:"column:usuario_id"`
	TotalAcesso uint          `gorm:"column:total_acesso"`
}

func (m AtivoMaisAcessadoModel) TableName() string {
	return "ativos_mais_acessados"
}

func (m AtivoMaisAcessadoModel) ToEntity() *entity.AtivoMaisAcessado {
	ativo := entity.NewAtivo(m.AtivoId)
	if m.Ativo != nil {
		ativo = m.Ativo.ToEntity()
	}

	usuario := entity.NewUsuario(m.UsuarioId)
	if m.Usuario != nil {
		usuario = m.Usuario.ToEntity()
	}

	return entity.NewAtivoMaisAcessado(
		m.Id,
		ativo,
		usuario,
		m.TotalAcesso,
	)
}
