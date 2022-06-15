package models

import "gorm.io/gorm"

type MySQL struct {
	Conexao *gorm.DB
}

func (p MySQL) Create(usr User) error {
	v := Macapa{Nome: usr.NameFormat(), Celular: usr.NumberFormat()}
	result := p.Conexao.Create(&v)
	return result.Error
}
