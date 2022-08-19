package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type MySQL struct {
	Conexao *gorm.DB
}

func (m MySQL) Create(usr User) error {
	v := Macapa{Nome: usr.NameFormat(), Celular: usr.NumberFormat()}
	result := m.Conexao.Create(&v)
	return result.Error
}

func (m MySQL) Find() ([]User, error) {
	var all []Macapa
	result := m.Conexao.Find(&all)

	// Go nao permite tipagem na heranca
	var users []User
	for _, usr := range all {
		users = append(users, usr)
	}
	return users, result.Error
}

func (m MySQL) Delete(nome string) error {
	result := m.Conexao.Where("nome LIKE ?", nome).Delete(&Macapa{})
	return result.Error
}

func (m MySQL) Update(id string, usr User) (User, error) {
	newUser := Macapa{Nome: usr.NameFormat(), Celular: usr.NumberFormat()}
	result := m.Conexao.Model(Varejao{}).Where("id LIKE ?", id).Updates(newUser.Format())
	return newUser, result.Error
}

func (m MySQL) FindByUser(usuario Login) (Login, error) {

	var usr Login
	m.Conexao.Where(&usuario).Find(&usr)

	if strings.Compare(usr.Email, usuario.Email) == 0 {
		return usr, nil
	}
	return usr, errors.New("Usuario nao Encontrado")
}
