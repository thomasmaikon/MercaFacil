package models

type Banco interface {
	Create(usr User) error
	Find() ([]User, error)
	Delete(nome string) error
	Update(id string, usr User) (User, error)
}
