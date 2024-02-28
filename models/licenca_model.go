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
	TipoId         uint              `gorm:"column:tipo_id"`
	Tipo           *LicencaTipoModel `gorm:"foreignKey:tipo_id"`
	CamadaId       *uint             `gorm:"column:camada_id"`
	Camada         *CamadaModel      `gorm:"foreignKey:camada_id"`
	AtivoId        uint              `gorm:"column:ativo_id"`
	Ativo          *AtivoModel       `gorm:"foreignKey:ativo_id"`
}

func (m *LicencaModel) TableName() string {
	return "licencas"
}

func (m *LicencaModel) ToEntity() *entity.Licenca {
	var dtProtocolo *time.Time
	var dtVencimento *time.Time

	if m.DataProtocolo != nil {
		_dtProtocolo, _ := time.Parse("2006-01-02 15:04:05", string(m.DataProtocolo))
		dtProtocolo = &_dtProtocolo
	}

	if m.DataVencimento != nil {
		_dtVencimento, _ := time.Parse("2006-01-02 15:04:05", string(m.DataVencimento))
		dtVencimento = &_dtVencimento
	}

	tipo := entity.NewLicencaTipo(m.TipoId)
	if m.Tipo != nil {
		tipo = m.Tipo.ToEntity()
	}

	var camada *entity.Camada
	if m.CamadaId != nil {
		camada = entity.NewCamada(*m.CamadaId)
	}

	if m.Camada != nil {
		camada = m.Camada.ToEntity()
	}

	ativo := entity.NewAtivo(m.AtivoId)
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
