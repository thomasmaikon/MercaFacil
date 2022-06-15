package models

type Usuario struct {
	Nome    string `json:"name"`
	Celular string `json:"cellphone"`
}

type ListUsuarios struct {
	Usrs []Usuario `json:"contacts"`
}

func (u Usuario) NameFormat() string {
	return u.Nome
}

func (u Usuario) NumberFormat() string {
	return u.Celular
}
