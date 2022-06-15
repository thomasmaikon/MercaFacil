package models

import "gorm.io/gorm"

type Postgres struct {
	Conexao *gorm.DB
}

func (p Postgres) Create(usr User) error {
	v := Varejao{Nome: usr.NameFormat(), Celular: usr.NumberFormat()}
	result := p.Conexao.Create(&v)
	return result.Error
}
