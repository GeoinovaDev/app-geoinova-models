package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type IncendioModel struct {
	Id               uint
	CamadaId         uint `gorm:"column:camada_id"`
	IncendioConfigId uint `gorm:"column:incendio_configuracao_id"`
	Wkt              string
	AreaHa           float64 `gorm:"column:area_ha"`
	Data             string
	Hora             string
	Confianca        string
	Criticidade      string
	Descricao        string
	Temperatura      int
	QuantidadeFocos  uint `gorm:"column:quantidade_focos"`
	Titulo           string
	Satelites        string
	Camada           *CamadaModel `gorm:"foreignKey:camada_id"`
}

func (m *IncendioModel) TableName() string {
	return "incendios_focos"
}

func (m *IncendioModel) ToEntity() *entity.Incendio {
	camada := entity.NewCamada(m.CamadaId)
	if m.Camada != nil {
		camada = m.Camada.ToEntity()
	}

	dataHoraStr := fmt.Sprintf("%s %s", m.Data, m.Hora)
	dataHora, err := time.Parse("2006-01-02 15:04:05", dataHoraStr)
	if err != nil {
		println(err)
	}

	return entity.NewIncendioBuilder(m.Id).
		WithWkt(m.Wkt).
		WithCamada(camada).
		WithAreaHa(m.AreaHa).
		WithCriticidade(m.Criticidade).
		WithConfianca(m.Confianca).
		WithDescricao(m.Descricao).
		WithTotalFocos(m.QuantidadeFocos).
		WithDateTime(&dataHora).
		WithTemperatura(m.Temperatura).
		WithTitulo(m.Titulo).
		WithSatelites(strings.Split(m.Satelites, ",")).
		Build()
}
