package models

import (
	"encoding/json"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type CamadaAtributoModel struct {
	Id       uint
	Json     string
	CamadaId uint `gorm:"column:camada_id"`
}

func (m *CamadaAtributoModel) TableName() string {
	return "camadas_atributos"
}

func (m *CamadaAtributoModel) ToEntity() []*entity.CamadaAtributo {
	atributos := []*entity.CamadaAtributo{}
	json.Unmarshal([]byte(m.Json), &atributos)
	return atributos
}
