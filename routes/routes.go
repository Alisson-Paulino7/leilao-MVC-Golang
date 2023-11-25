package routes

import (
	"goapp/controllers"
	"net/http"
)

// Carrega as rotas da aplicação
func CarregaRotas() {
	http.HandleFunc("/login", controllers.ValidarLogin)
	http.HandleFunc("/cadastro", controllers.CadastrarUser)
	http.HandleFunc("/listar", controllers.ListarProdutos)
	http.HandleFunc("/SalvarProduto", controllers.SalvarProduto)
}
