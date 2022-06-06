package models

const (
	TipoMacapa  string = "1"
	TipoVarejao string = "2"
)

type Login struct {
	Email string `json:"email" gorm:" unique; not null; varchar(100)"`
	Senha string `json:"senha" gorm:" not null; varchar(100)"`
	Tipo  string `gorm:" not null"`
}
