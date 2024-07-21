package routes

import (
	"lojaGo/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)  /*Função com a rota que especifica quem vai atender. Aqui indica o caminho da página "Index"*/
	http.HandleFunc("/new", controllers.New) /*Função com a rota que especifica quem vai atender. Aqui indica o caminho da página "New"*/
}
