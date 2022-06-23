package models

import "gorm.io/gorm"

type Postgres struct {
	Conexao *gorm.DB
}

func (p Postgres) Create(usr User) error {
	v := Varejao{Nome: usr.NameFormat(), Celular: usr.NumberFormat()}
	result := p.Conexao.Create(v.Format())
	return result.Error
}

func (p Postgres) Find() ([]User, error) {
	var all []Varejao
	result := p.Conexao.Find(&all)

	// Go nao permite tipagem na heranca
	var users []User
	for _, usr := range all {
		users = append(users, usr)
	}
	return users, result.Error
}

func (p Postgres) Delete(nome string) error {
	result := p.Conexao.Where("nome LIKE ?", nome).Delete(&Varejao{})
	return result.Error
}

func (p Postgres) Update(id string, usr User) (User, error) {
	newUser := Varejao{Nome: usr.NameFormat(), Celular: usr.NumberFormat()}
	result := p.Conexao.Model(Varejao{}).Where("id = ?", id).Updates(newUser.Format())
	return newUser, result.Error
}
