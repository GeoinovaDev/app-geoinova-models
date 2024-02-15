package models

type FormularioPropriedadeAtributo struct {
	Id            uint
	Atributo      string
	PropriedadeId uint `gorm:"column:propriedade_id"`
}

func (FormularioPropriedadeAtributo) TableName() string {
	return "formularios_propriedades_atributos"
}
