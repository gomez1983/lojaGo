package models

import "lojaGo/db"

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados() /*Abre a conexão com o BD*/

	selectDeTodosOsProdutos, err := db.Query("select * from produtos") /*Faz o select com o DB para trazer os produtos*/
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error()) /*A função panic é usada para interromper a execução do programa*/
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
