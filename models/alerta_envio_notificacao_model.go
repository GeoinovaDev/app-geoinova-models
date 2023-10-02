package models

import (
	"encoding/json"
	"strings"

	"github.com/GeoinovaDev/app-geoinova-entity/entity"
)

type AlertaEnvioNotificacaoModel struct {
	Id                 uint
	Uuid               string
	AtivoNome          string `gorm:"column:ativo_nome"`
	DataImagemAntes    string `gorm:"data_imagem_antes"`
	DataImagemDepois   string `gorm:"data_imagem_depois"`
	DeteccaoClasseNome string `gorm:"deteccao_classe_nome"`
	DeteccaoDescricao  string `gorm:"deteccao_descricao"`
	AnaliseAutomatica  bool   `gorm:"analise_automatica"`
	Posicao            uint
	Emails             string
	Telefones          string
	Parametros         string
}

func (m AlertaEnvioNotificacaoModel) TableName() string {
	return "alertas_envio_notificacoes"
}

func (m AlertaEnvioNotificacaoModel) ToEntity() *entity.AlertaEnvioNotificacao {
	emails := []*entity.EnderecoEmail{}
	telefones := []*entity.Telefone{}

	_emails := strings.Split(m.Emails, ";")
	for _, email := range _emails {
		emails = append(emails, entity.NewEnderecoEmail("", email))
	}

	_telefones := strings.Split(m.Telefones, ";")
	for _, telefone := range _telefones {
		telefones = append(telefones, entity.ParseTelefone(telefone))
	}

	parametros := make(map[string]string)
	json.Unmarshal([]byte(m.Parametros), &parametros)

	return entity.
		NewAlertaNotificacaoBuilder().
		WithEmails(emails).
		WithAtivoNome(m.AtivoNome).
		WithDataImagemAntes(m.DataImagemAntes).
		WithDataImagemDepois(m.DataImagemDepois).
		WithDeteccaoClasseNome(m.DeteccaoClasseNome).
		WithDeteccaoDescricao(m.DeteccaoDescricao).
		WithAnaliseAutomatica(m.AnaliseAutomatica).
		WithPosicao(m.Posicao).
		WithTelefones(telefones).
		WithParametros(parametros).
		Build()
}
