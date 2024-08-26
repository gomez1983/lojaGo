package routes

import (
	"lojaGo/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)        /*Função com a rota que especifica que vai atender. Aqui indica o caminho da página "Index"*/
	http.HandleFunc("/new", controllers.New)       /*Função com a rota que especifica que vai atender. Aqui indica o caminho da página "New"*/
	http.HandleFunc("/insert", controllers.Insert) /*Função com a rota que especifica que vai atender. Aqui indica o caminho quando tiver requisição para Insert*/
	http.HandleFunc("/delete", controllers.Delete) /*Função com a rota que especifica que vai atender. Aqui indica o caminho quando tiver requisição para Deletar*/
	http.HandleFunc("/edit", controllers.Edit)     /*Função com a rota que especifica que vai atender. Aqui indica o caminho quando tiver requisição para Editar*/
	http.HandleFunc("/update", controllers.Update) /*Função com a rota que especifica que vai atender. Aqui indica o caminho quando tiver requisição para Atualizar o que for editado*/
}
