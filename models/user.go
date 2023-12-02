// models/models.go
package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

// Estrutura que fiz para o usuario

type Usuario struct {
	ID    int    	`json:"id"`
	Email string 	`json:"email"`
	Senha string 	`json:"senha"`
	Cpf string 		`json:"cpf"`
	Telefone string `json:"telefone"`
	Foto[]byte 		`json:"foto"`
}

// Constante pra facilitar e não precisar fazer
// manualmente toda vez que precisa se conetar com o mysql

const (
	dsn = "go_app_user:go_app_password@tcp(mysql:3306)/go_app_db"
)

// Método pra inserir os dados da estrutura no mysql

func (u *Usuario) AdicionaUsuario() error {
	//gerando conexão com meu banco e passando os parâmetros
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar hash da senha: %v", err)
	}

	result, err := db.Exec("INSERT INTO usuarios (email, senha, cpf, telefone) VALUES (?, ?, ?, ?)",
		u.Email, hashedSenha, u.Cpf, u.Telefone)

	if err != nil {
		return fmt.Errorf("erro ao inserir usuário no banco de dados: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter número de linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma linha foi afetada pela inserção")
	}

	return nil
}

func (u *Usuario) ValidarLogin() error {
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	var senha string
	err = db.QueryRow("SELECT id, senha FROM usuarios WHERE email = ?", u.Email).Scan(&senha)
	if err != nil {
		return fmt.Errorf("erro ao buscar usuário no banco de dados: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(senha), []byte(u.Senha))
	if err != nil {
		return fmt.Errorf("senha inválida: %v", err)
	}
	
	return nil
}