package db

import (
	"database/sql"

	_ "github.com/lib/pq" /*O underline serve para manter o import mesmo se ele não for utilizado*/
)

func ConectaComBancoDeDados() *sql.DB { /*Quando importada uma biblioteca, a função tem que ter a primeira letra em caixa alta*/
	conexao := "user=postgres dbname=alura_loja password=catita01 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
