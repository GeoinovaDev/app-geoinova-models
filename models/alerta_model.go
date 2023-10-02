package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type AlertaModel struct {
	Id            uint
	Uuid          string
	Observacao    string
	Responsavel   string
	Justificativa string
	Status        string
	Posicao       uint           `gorm:"column:posicao"`
	Deteccao      *DeteccaoModel `gorm:"foreignKey:deteccao_id"`
	DeteccaoId    uint           `gorm:"column:deteccao_id"`
	AnaliseId     uint           `gorm:"column:analise_id"`
	Analise       *AnaliseModel  `gorm:"foreignKey:analise_id"`
}

func (m AlertaModel) TableName() string {
	return "alertas"
}

func (m AlertaModel) ToEntity() *entity.Alerta {
	deteccao := entity.NewDeteccaoBuilder(m.DeteccaoId).Build()
	if m.Deteccao != nil {
		deteccao = m.Deteccao.ToEntity()
	}

	analise := entity.NewAnalise(m.AnaliseId)
	if m.Analise != nil {
		analise = m.Analise.ToEntity()
	}

	return entity.
		NewAlertaBuilder(m.Id).
		WithPosicao(m.Posicao).
		WithObservacao(m.Observacao).
		WithResponsavel(m.Responsavel).
		WithJustificativa(m.Justificativa).
		WithStatus(entity.AlertaFaseStatus(m.Status)).
		WithDeteccao(deteccao).
		WithAnalise(analise).
		Build()
}
