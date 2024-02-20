package models

import (
	"strings"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type AppPreferModel struct {
	Id                   uint
	ColunasTabelaCamadas string `gorm:"column:colunas_tabela_camadas"`
}

func (m *AppPreferModel) TableName() string {
	return "app_prefer"
}

func (m *AppPreferModel) ToEntity() *entity.AppPrefer {
	return entity.
		NewAppPreferBuilder(m.Id).
		WithColunasTabelaCamadas(strings.Split(m.ColunasTabelaCamadas, ";")).
		Build()
}
