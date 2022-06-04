package models

type Login struct {
	Email string `json:"email" gorm:" unique; not null; varchar(100)"`
	Senha string `json:"senha" gorm:" not null; varchar(100)"`
	Tipo  int    `gorm:" not null"`
}
