package routes

import (
	"lojaGo/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) /*Função com a rota que especifica quem vai atender*/
}
