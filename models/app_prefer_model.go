package models

import (
	"strings"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type AppPreferModel struct {
	Id                   uint
	ClienteId            uint   `gorm:"column:cliente_id"`
	ColunasTabelaCamadas string `gorm:"column:colunas_tabela_camadas"`
}

func (m *AppPreferModel) TableName() string {
	return "app_prefers"
}

func (m *AppPreferModel) ToEntity() *entity.AppPrefer {
	return entity.
		NewAppPreferBuilder(m.Id).
		WithCliente(entity.NewCliente(m.ClienteId)).
		WithColunasTabelaCamadas(strings.Split(m.ColunasTabelaCamadas, ";")).
		Build()
}
