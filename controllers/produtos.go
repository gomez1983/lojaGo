package controllers

import (
	"lojaGo/models"
	"net/http"
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
