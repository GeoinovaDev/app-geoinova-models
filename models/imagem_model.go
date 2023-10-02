package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type ImagemModel struct {
	Id                 uint
	Uuid               string
	Data               string
	Bucket             string
	Resolucao          string
	Provedor           string
	TotalBandas        uint             `gorm:"column:total_bandas"`
	TaxaCoberturaNuvem float32          `gorm:"column:taxa_cobertura_nuvem"`
	ProvedorFonte      string           `gorm:"column:provedor_fonte"`
	AtivoId            uint             `gorm:"column:ativo_id"`
	Ativo              *AtivoModel      `gorm:"foreignKey:ativo_id"`
	QualidadeId        uint             `gorm:"column:qualidade_id"`
	Qualidade          *ImagemQualidade `gorm:"foreignKey:qualidade_id"`
	Source             int
}

func (m ImagemModel) TableName() string {
	return "imagens"
}

func (m ImagemModel) ToEntity() *entity.Imagem {
	qualidade := entity.NewImagemQualidade(m.QualidadeId)

	if m.Qualidade != nil {
		qualidade = m.Qualidade.ToEntity()
	}

	ativo := entity.NewAtivo(m.AtivoId)
	if m.Ativo != nil {
		ativo = m.Ativo.ToEntity()
	}

	return entity.
		NewImagemBuilder(m.Id).
		WithUuid(m.Uuid).
		WithData(m.Data).
		WithBucket(m.Bucket).
		WithResolucao(m.Resolucao).
		WithTotalBanda(m.TotalBandas).
		WithTaxaCoberturaNuvem(m.TaxaCoberturaNuvem).
		WithProvedor(m.Provedor).
		WithProvedorFonte(m.ProvedorFonte).
		WithQualidade(qualidade).
		WithAtivo(ativo).
		WithSource(entity.ImagemSource(m.Source)).
		Build()
}
