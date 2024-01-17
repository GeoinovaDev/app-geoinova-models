package models

import "github.com/GeoinovaDev/app-geoinova-entity/entity"

type AtivoBucketModel struct {
	Id        uint
	AtivoId   uint        `gorm:"column:ativo_id"`
	Ativo     *AtivoModel `gorm:"foreignKey:ativo_id"`
	Bucket    string
	IsDefault bool `gorm:"column:is_default"`
}

func (m *AtivoBucketModel) TableName() string {
	return "ativos_buckets"
}

func (m *AtivoBucketModel) ToEntity() *entity.AtivoBucket {
	return entity.
		NewAtivoBucketBuilder(m.Id).
		WithAtivo(entity.NewAtivo(m.AtivoId)).
		WithBucket(m.Bucket).
		WithIsDefault(m.IsDefault).
		Build()
}
