package models

import (
	"time"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type AlertaEventoModel struct {
	Id        uint
	Status    string
	Param     string
	AlertaId  uint         `gorm:"column:alerta_id"`
	Alerta    *AlertaModel `gorm:"foreignKey:alerta_id"`
	CreatedAt []uint8      `gorm:"column:created_at"`
}

func (m AlertaEventoModel) TableName() string {
	return "alertas_eventos"
}

func (m AlertaEventoModel) ToEntity() *entity.AlertaEvento {
	alerta := entity.NewAlerta(m.AlertaId)
	if m.Alerta != nil {
		alerta = m.Alerta.ToEntity()
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05", string(m.CreatedAt))
	if err != nil {
	}

	return entity.
		NewAlertaEventoBuilder(m.Id, entity.AlertaEventoStatus(m.Status)).
		WithParam(m.Param).
		WithAlerta(alerta).
		WithCreatedAt(createdAt).
		Build()
}
