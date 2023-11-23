package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"html/template"
	// "log"

	"goapp/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsn = "go_app_user:go_app_password@tcp(mysql:3306)/go_app_db"
)

func CadastrarUser(w http.ResponseWriter, r *http.Request) {

	// Validando se o método está como post

	if r.Method == "POST" {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		// 	return
		// }
		// Validando se o Mysql está de boa pra receber conexão
		err := waitForMySQL()
		if err != nil {
			http.Error(w, "Erro ao aguardar o MySQL iniciar", http.StatusInternalServerError)
			return
		}
		// Pegando os dados do Forms lá do html
		email := r.FormValue("email")
		senha := r.FormValue("senha")
		cpf := r.FormValue("cpf")
		telefone := r.FormValue("telefone")

		if email == "" || senha == "" || cpf == "" || telefone == "" {
			http.Error(w, "Por favor, preencha todos os campos", http.StatusBadRequest)
			return
		}
		// Instanciano o modelo da estrutura Usuario e passando ps parametros recebidos no Post
		usuario := models.Usuario{
			Email:    email,
			Senha:    senha,
			Cpf:      cpf,
			Telefone: telefone,
		}
		// Método do models pra inserir os dados e tratar algum erro que tiver
		err = usuario.AdicionaUsuario()
		// checaErro(err)
		if err != nil {
			http.Error(w, "Erro ao cadastrar usuário no banco de dados", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ValidarLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := waitForMySQL()
		if err != nil {
			http.Error(w, "Erro ao aguardar o MySQL iniciar", http.StatusInternalServerError)
			return
		}

		email := r.FormValue("email")
		senha := r.FormValue("senha")

		if email == "" || senha == "" {
			http.Error(w, "Por favor, preencha todos os campos", http.StatusBadRequest)
			return
		}

		usuario := models.Usuario{
			Email: email,
			Senha: senha,
		}
		err = usuario.ValidarLogin()
		if err != nil {
			http.Error(w, "Erro ao validar usuário no banco de dados", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "cadastro.html", http.StatusSeeOther)
	}
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {

	err := waitForMySQL()
	if err != nil {
		http.Error(w, "Erro ao aguardar o MySQL iniciar", http.StatusInternalServerError)
		return
	}

	produtos, err := models.ListarProdutos()
	if err != nil {
		// Verificando erro em situação de erro no produto ou consulta zerada
		http.Error(w, fmt.Sprintf("Erro ao listar produtos: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("app/templates/listaProdutos.html")
	if err != nil {
		http.Error(w, "Erro ao analisar o modelo HTML", http.StatusInternalServerError)
		return
	}

	// tmpl, err := template.New("listaProdutos").Parse(`
	// <!DOCTYPE html>
	// <html>
	// <head>
	// 	<title>Lista de Produtos</title>
	// </head>
	// <body>
	// 	<h1>Lista de Produtos</h1>
	// 	<ul>
	// 		{{range .}}
	// 			<li>
	// 				<img src="data:image/png;base64,{{.FotoProduct}}" alt="Foto do Produto">
	// 				<strong>{{.Nome}}</strong> - 
	// 				R$ {{.LanceInicial}} - 
	// 				Expira em: {{.TmpExpiracao}} dias
	// 			</li>
	// 		{{end}}
	// 	</ul>
	// </body>
	// </html>
	// `)
	
if err != nil {
	http.Error(w, "Erro ao analisar o modelo HTML", http.StatusInternalServerError)
	return
}

err = tmpl.Execute(w, produtos)
if err != nil {
	http.Error(w, "Erro ao executar o template HTML", http.StatusInternalServerError)
	return
}



	

}

func waitForMySQL() error {
	// Fica tentando a conexão por 30 segundos
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("tempo limite atingido ao aguardar o MySQL iniciar")
		case <-ticker.C:
			db, err := sql.Open("mysql", dsn)
			// checaErro(err)
			if err == nil {
				defer db.Close()

				// faz uma consulta simples para verificar se o MySQL tá pronto
				err = db.Ping()
				// checaErro(err)
				if err == nil {
					return nil
				}
			}
		}
	}
}

// func checaErro(err error) {
// 	if err != nil {
// 		log.Fatal("Erro encontrado ao processar pedido: ", err)
// 	}
// }
