package models

import (
	"strconv"
	"strings"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type DeteccaoModel struct {
	Id                   uint
	Wkt                  string
	Descricao            string
	AreaHa               float64              `gorm:"column:area_ha"`
	PreviewAntesUuid     string               `gorm:"preview_antes_uuid"`
	PreviewDepoisUuid    string               `gorm:"preview_depois_uuid"`
	PreviewResultadoUuid string               `gorm:"preview_resultado_uuid"`
	ClasseId             uint                 `gorm:"column:classe_id"`
	ClienteId            uint                 `gorm:"column:cliente_id"`
	Classe               *DeteccaoClasseModel `gorm:"foreignKey:classe_id"`
	PreviewAntesNome     string               `gorm:"preview_antes_nome"`
	PreviewDepoisNome    string               `gorm:"preview_depois_nome"`
	PreviewRegiaoNome    string               `gorm:"preview_regiao_nome"`
	BoundsAntes          string               `gorm:"bounds_antes"`
	BoundsDepois         string               `gorm:"bounds_depois"`
	BoundsRegiao         string               `gorm:"bounds_regiao"`
}

func (m DeteccaoModel) TableName() string {
	return "deteccoes"
}

func (m DeteccaoModel) ToEntity() *entity.Deteccao {
	classe := entity.NewDeteccaoClasseBuilder(m.ClasseId).Build()
	if m.Classe != nil {
		classe = m.Classe.ToEntity()
	}

	_boundsAntes := strings.Split(m.BoundsAntes, ";")
	var boundsAntes []float32
	for _, b := range _boundsAntes {
		bound, _ := strconv.ParseFloat(b, 32)
		boundsAntes = append(boundsAntes, float32(bound))
	}

	_boundsDepois := strings.Split(m.BoundsDepois, ";")
	var boundsDepois []float32
	for _, b := range _boundsDepois {
		bound, _ := strconv.ParseFloat(b, 32)
		boundsDepois = append(boundsDepois, float32(bound))
	}

	_boundsRegiao := strings.Split(m.BoundsRegiao, ";")
	var boundsRegiao []float32
	for _, b := range _boundsRegiao {
		bound, _ := strconv.ParseFloat(b, 32)
		boundsRegiao = append(boundsRegiao, float32(bound))
	}

	return entity.
		NewDeteccaoBuilder(m.Id).
		WithWkt(m.Wkt).
		WithDescricao(m.Descricao).
		WithImagemAntesUuid(m.PreviewAntesUuid).
		WithImagemDepoisUuid(m.PreviewDepoisUuid).
		WithImagemResultadoUuid(m.PreviewResultadoUuid).
		WithCliente(entity.NewCliente(m.ClienteId)).
		WithClasse(classe).
		WithPreviews(entity.
			NewDeteccaoPreviewsBuilder().
			WithAntes(entity.
				NewDeteccaoPreviewBuilder().
				WithNome(m.PreviewAntesNome).
				Build()).
			WithDepois(entity.
				NewDeteccaoPreviewBuilder().
				WithNome(m.PreviewDepoisNome).
				Build()).
			WithRegiao(entity.
				NewDeteccaoPreviewBuilder().
				WithNome(m.PreviewRegiaoNome).
				Build()).
			Build()).
		WithBounds(entity.
			NewDeteccaoBoundsBuilder().
			WithAntes(boundsAntes).
			WithDepois(boundsDepois).
			WithRegiao(boundsRegiao).
			Build()).
		Build()
}
