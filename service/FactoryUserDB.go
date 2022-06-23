package service

import (
	"errors"
	"thomas/projeto_mercafacil/db"
	m "thomas/projeto_mercafacil/models"
)

type FactoryUser struct {
}

func (f FactoryUser) GetUserDB(tipo string) (m.Banco, error) {

	if tipo == m.TipoVarejao {

		return m.Postgres{Conexao: db.GetPostgresConnection()}, nil
	}

	if tipo == m.TipoMacapa {
		return m.MySQL{Conexao: db.GetMysqlConnection()}, nil
	}

	return nil, errors.New("Falha ao fazer a identificacao do usuario")
}
