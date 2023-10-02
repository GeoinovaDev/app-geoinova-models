package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type ImagemFavoritadaModel struct {
	Id        uint
	UsuarioId uint          `gorm:"column:usuario_id"`
	Usuario   *UsuarioModel `gorm:"foreignKey:usuario_id"`
	ImagemId  uint          `gorm:"column:imagem_id"`
	Imagem    *ImagemModel  `gorm:"foreignKey:imagem_id"`
}

func (m ImagemFavoritadaModel) TableName() string {
	return "imagens_favoritadas"
}

func (m ImagemFavoritadaModel) ToEntity() *entity.ImagemFavoritada {
	usuario := entity.NewUsuario(m.UsuarioId)
	if m.Usuario != nil {
		usuario = m.Usuario.ToEntity()
	}

	imagem := entity.NewImagem(m.ImagemId)
	if m.Imagem != nil {
		imagem = m.Imagem.ToEntity()
	}

	return entity.NewImagemFavoritada(
		m.Id,
		usuario,
		imagem,
	)
}
