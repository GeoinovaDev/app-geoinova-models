package models

type IncendioCamada struct {
	Id               uint
	CamadaId         uint `gorm:"column:camada_id"`
	IncendioConfigId uint `gorm:"column:incendio_configuracao_id"`
}

func (m *IncendioCamada) TableName() string {
	return "incendios_camadas"
}
