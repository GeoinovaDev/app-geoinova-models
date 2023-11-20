package models

import (
	"time"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type AlertaFaseModel struct {
	Id        uint
	Status    string
	AlertaId  uint         `gorm:"column:alerta_id"`
	Alerta    *AlertaModel `gorm:"foreignKey:alerta_id"`
	CreatedAt []uint8      `gorm:"column:created_at"`
}

func (m AlertaFaseModel) TableName() string {
	return "alertas_fases"
}

func (m AlertaFaseModel) ToEntity() *entity.AlertaFase {
	alerta := entity.NewAlerta(m.AlertaId)
	if m.Alerta != nil {
		alerta = m.Alerta.ToEntity()
	}

	createdAt, _ := time.Parse("01-02-2006 15:04:05", string(m.CreatedAt))
	return entity.
		NewAlertaFaseBuilder(m.Id, entity.AlertaFaseStatus(m.Status)).
		WithAlerta(alerta).
		WithCreatedAt(createdAt).
		Build()
}
