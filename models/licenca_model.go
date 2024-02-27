package models

import (
	"time"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type LicencaModel struct {
	Id             uint
	Nome           string
	Descricao      string
	Status         string
	Filename       string            `gorm:"column:file_name"`
	FileBucket     string            `gorm:"column:file_bucket"`
	FileSource     string            `gorm:"column:file_source"`
	DataProtocolo  []uint8           `gorm:"column:data_protocolo"`
	DataVencimento []uint8           `gorm:"column:data_vencimento"`
	TipoId         uint              `gorm:"column:tipoca_id"`
	Tipo           *LicencaTipoModel `gorm:"foreignKey:tipoca_id"`
	CamadaId       uint              `gorm:"column:camada_id"`
	Camada         *CamadaModel      `gorm:"foreignKey:camada_id"`
	AtivoId        uint              `gorm:"column:ativo_id"`
	Ativo          *AtivoModel       `gorm:"foreignKey:ativo_id"`
}

func (m *LicencaModel) TableName() string {
	return "licencas"
}

func (m *LicencaModel) ToEntity() *entity.Licenca {
	dtProtocolo, _ := time.Parse("2006-01-02 15:04:05", string(m.DataProtocolo))
	dtVencimento, _ := time.Parse("2006-01-02 15:04:05", string(m.DataVencimento))

	var tipo *entity.LicencaTipo
	if m.Tipo != nil {
		tipo = m.Tipo.ToEntity()
	}

	var camada *entity.Camada
	if m.Camada != nil {
		camada = m.Camada.ToEntity()
	}

	var ativo *entity.Ativo
	if m.Ativo != nil {
		ativo = m.Ativo.ToEntity()
	}

	return entity.
		NewLicencaBuilder(m.Id).
		WithNome(m.Nome).
		WithDescricao(m.Descricao).
		WithStatus(entity.LicencaStatus(m.Status)).
		WithArquivo(entity.
			NewArquivoBuilder(entity.NewArquivoFromFilename(m.Filename)).
			WithSource(m.FileSource).
			WithBucket(m.FileBucket).
			Build()).
		WithDataProtocolo(dtProtocolo).
		WithDataVencimento(dtVencimento).
		WithTipo(tipo).
		WithCamada(camada).
		WithAtivo(ativo).
		Build()
}
