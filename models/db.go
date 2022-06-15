package models

type Banco interface {
	Create(usr User) error
}
