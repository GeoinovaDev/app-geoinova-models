package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type DeteccaoCamadaModel struct {
	Id         uint
	Deteccao   *DeteccaoModel `gorm:"foreignKey:deteccao_id"`
	DeteccaoId uint           `gorm:"column:deteccao_id"`
	Camada     *CamadaModel   `gorm:"foreignKey:camada_id"`
	CamadaId   uint           `gorm:"column:camada_id"`
}

func (m DeteccaoCamadaModel) TableName() string {
	return "deteccoes_camadas"
}

func (m DeteccaoCamadaModel) ToEntity() *entity.DeteccaoCamada {
	deteccao := entity.NewDeteccao(m.DeteccaoId)
	if m.Deteccao != nil {
		deteccao = m.Deteccao.ToEntity()
	}

	camada := entity.NewCamada(m.CamadaId)
	if m.Camada != nil {
		camada = m.Camada.ToEntity()
	}

	return entity.NewDeteccaoCamada(
		m.Id,
		deteccao,
		camada,
	)
}
