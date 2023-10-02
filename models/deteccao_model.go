package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type DeteccaoModel struct {
	Id                   uint
	Wkt                  string
	Descricao            string
	PreviewAntesUuid     string               `gorm:"preview_antes_uuid"`
	PreviewDepoisUuid    string               `gorm:"preview_depois_uuid"`
	PreviewResultadoUuid string               `gorm:"preview_resultado_uuid"`
	ClasseId             uint                 `gorm:"column:classe_id"`
	Classe               *DeteccaoClasseModel `gorm:"foreignKey:classe_id"`
}

func (m DeteccaoModel) TableName() string {
	return "deteccoes"
}

func (m DeteccaoModel) ToEntity() *entity.Deteccao {
	classe := entity.NewDeteccaoClasseBuilder(m.ClasseId).Build()
	if m.Classe != nil {
		classe = m.Classe.ToEntity()
	}

	return entity.
		NewDeteccaoBuilder(m.Id).
		WithWkt(m.Wkt).
		WithDescricao(m.Descricao).
		WithImagemAntesUuid(m.PreviewAntesUuid).
		WithImagemDepoisUuid(m.PreviewDepoisUuid).
		WithImagemResultadoUuid(m.PreviewResultadoUuid).
		WithClasse(classe).
		Build()
}
