package service

import (
	"errors"
	"thomas/projeto_mercafacil/db"
	m "thomas/projeto_mercafacil/models"
)

type FactoryUser struct {
}

func (f FactoryUser) GetUserDB(tipo string, usuarios []m.Usuario) ([]m.User, m.Banco, error) {

	if tipo == m.TipoVarejao {
		list := []m.User{}
		for _, usr := range usuarios {
			list = append(list, m.Varejao{Nome: usr.Nome, Celular: usr.Celular})
		}
		return list, m.Postgres{Conexao: db.GetPostgresConnection()}, nil
	}

	if tipo == m.TipoMacapa {
		list := []m.User{}
		for _, usr := range usuarios {
			list = append(list, m.Macapa{Nome: usr.Nome, Celular: usr.Celular})
		}
		return list, m.MySQL{Conexao: db.GetMysqlConnection()}, nil
	}

	return nil, nil, errors.New("Falha ao fazer a identificacao do usuario")
}
