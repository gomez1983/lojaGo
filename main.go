package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq" /*O underline serve para manter o import mesmo se ele não for utilizado*/
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=catita01 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) /* Encapsula todos os templates html do pacote para a variável*/

func main() {
	http.HandleFunc("/", index) /*Função com a rota que especifica quem vai atender*/
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) { /*Index é quem vai atender a rota*/
	db := conectaComBancoDeDados() /*Abre a conexão com o BD*/

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

	temp.ExecuteTemplate(w, "Index", produtos) /*Cria a lista de produtos e repassa ela para o template em Index*/
	defer db.Close()
}
