package models

import (
	"encoding/json"
)

type CamadaAtributoModel struct {
	Id       uint
	Json     string
	CamadaId uint `gorm:"column:camada_id"`
}

func (m *CamadaAtributoModel) TableName() string {
	return "camadas_atributos"
}

func (m *CamadaAtributoModel) ToEntity() map[string]string {
	atributos := map[string]string{}
	json.Unmarshal([]byte(m.Json), &atributos)
	return atributos
}
