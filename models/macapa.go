package models

import (
	"strings"
)

type Macapa struct {
	Id      int    `gorm:"primaryKey; autoincrement"`
	Nome    string `json:"name" gorm:"unique; not null; varchar(200)"`
	Celular string `json:"cellphone" gorm:"not null; varchar(20)"`
}

type ListMacapaUsers struct {
	Usrs []Macapa `json:"contacts"`
}

func (m Macapa) Format() *Macapa {
	return &Macapa{
		Nome:    m.NameFormat(),
		Celular: m.NumberFormat(),
	}
}

func (m Macapa) NameFormat() string {
	return strings.ToUpper(m.Nome)
}

func (m Macapa) NumberFormat() string {
	numeroLimpo := strings.Replace(m.Celular, "[^\\d.]", "", -1)
	pais := numeroLimpo[0:2]
	ddd := numeroLimpo[2:4]
	part1 := numeroLimpo[4:9]
	part2 := numeroLimpo[9:]

	return "+" + pais + " (" + ddd + ") " + part1 + "-" + part2
}
