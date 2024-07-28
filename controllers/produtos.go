package controllers

import (
	"log"
	"lojaGo/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) /* Encapsula todos os templates html do pacote para a variável*/

func Index(w http.ResponseWriter, r *http.Request) { /*Index é quem vai atender a rota*/
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) /*Cria a lista de produtos e repassa ela para o template em Index*/
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64) /*Converte a variável preço (que chega como String) para Float64*/
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade) /*Converte a variável quantidade (que chega como String) para Int*/
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301) /*Código 301 é responsável por sinalizar que deu tudo certo com a requisição */
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id") /*A função pega a URL do botão DELETE*/
	models.DeletaProduto(idDoProduto)      /*Deleta o produto*/
	http.Redirect(w, r, "/", 301)          /*W é o responseWriter, r é a requisição, '/' é a página principal e 301 é o código do retorno */
}
