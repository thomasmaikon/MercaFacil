package models

import "strings"

type Varejao struct {
	Id      int    `gorm:"primaryKey; autoincrement"`
	Nome    string `json:"name" gorm:"unique; not null; varchar(200)"`
	Celular string `json:"cellphone" gorm:"not null; varchar(13)"`
}

type ListVarejaoUsers struct {
	Usrs []Varejao `json:"contacts"`
}

func (v Varejao) Format() *Varejao {
	return &Varejao{
		Nome:    v.NameFormat(),
		Celular: v.NumberFormat(),
	}
}

func (v Varejao) NameFormat() string {
	return v.Nome
}

func (v Varejao) NumberFormat() string {
	return strings.Replace(v.Celular, "[^\\d.]", "", -1)
}
