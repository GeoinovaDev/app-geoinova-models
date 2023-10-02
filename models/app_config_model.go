package models

type AppConfigModel struct {
	Id    uint
	Name  string
	Value string
}

func (m AppConfigModel) TableName() string {
	return "app_config"
}
