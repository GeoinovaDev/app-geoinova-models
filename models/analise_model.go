package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type AnaliseModel struct {
	Id             uint
	Wkt            string
	Uuid           string
	Status         string
	Tipo           string
	ImagemAntesId  uint          `gorm:"column:imagem_antes_id"`
	ImagemAntes    *ImagemModel  `gorm:"foreignKey:imagem_antes_id"`
	ImagemDepoisId uint          `gorm:"column:imagem_depois_id"`
	ImagemDepois   *ImagemModel  `gorm:"foreignKey:imagem_depois_id"`
	TotalArea      float32       `gorm:"column:total_area"`
	AtivoId        uint          `gorm:"column:ativo_id"`
	Ativo          *AtivoModel   `gorm:"foreignKey:ativo_id"`
	UsuarioId      uint          `gorm:"column:usuario_id"`
	Usuario        *UsuarioModel `gorm:"foreignKey:usuario_id"`
}

func (m AnaliseModel) TableName() string {
	return "analises"
}

func (m AnaliseModel) ToEntity() *entity.Analise {
	antes := entity.NewImagem(m.ImagemAntesId)
	if m.ImagemAntes != nil {
		antes = m.ImagemAntes.ToEntity()
	}

	depois := entity.NewImagem(m.ImagemDepoisId)
	if m.ImagemDepois != nil {
		depois = m.ImagemDepois.ToEntity()
	}

	ativo := entity.NewAtivo(m.AtivoId)
	if m.Ativo != nil {
		ativo = m.Ativo.ToEntity()
	}

	usuario := entity.NewUsuario(m.UsuarioId)
	if m.Usuario != nil {
		usuario = m.Usuario.ToEntity()
	}

	return entity.
		NewAnaliseBuilder(m.Id).
		WithWkt(m.Wkt).
		WithStatus(entity.AnaliseStatus(m.Status)).
		WithTipo(entity.AnaliseTipo(m.Tipo)).
		WithTotalArea(m.TotalArea).
		WithAtivo(ativo).
		WithImagemAntes(antes).
		WithImagemDepois(depois).
		WithUsuario(usuario).
		Build()
}
