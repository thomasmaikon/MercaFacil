package service

import (
	"errors"
	"strings"
	"thomas/projeto_mercafacil/db"
	m "thomas/projeto_mercafacil/models"
)

type FactoryUser struct {
}

func (f FactoryUser) GetUserDB(tipo string) (m.Banco, error) {

	if strings.Compare(tipo, m.TipoVarejao) == 0 {

		return m.Postgres{Conexao: db.GetPostgresConnection()}, nil
	}

	if strings.Compare(tipo, m.TipoMacapa) == 0 {
		return m.MySQL{Conexao: db.GetMysqlConnection()}, nil
	}

	return nil, errors.New("Falha ao fazer a identificacao do usuario")
}
