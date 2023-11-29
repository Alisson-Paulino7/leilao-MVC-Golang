package models

import (
	"database/sql"
	"fmt"
	"encoding/base64"

	_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	ID           int   		`json:"id"`
	Nome         string 	`json:"nome"`
	Descricao	 string 	`json:"descricao"`
	LanceInicial int 		`json:"lance_inicial"`
	TmpExpiracao int 		`json:"tmp_expiracao"`
	FotoProduct  string 	`json:"Foto_Product"`
}

// ListarProdutos retorna uma lista de todos os produtos cadastrados.
func ListarProdutos() ([]Produto, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta: %v", err)
	}
	defer rows.Close()

	// Montar a estrutura com o que recebo do banco de dados
	var produtos []Produto
	for rows.Next() {
		var produto Produto
		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.LanceInicial, &produto.TmpExpiracao, &produto.FotoProduct)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler resultado: %v", err)
		}
		produto.FotoProduct = base64.StdEncoding.EncodeToString([]byte(produto.FotoProduct))

		produtos = append(produtos, produto)
	}

	// Verificar depois do loop se deu problema em algo
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("erro após o loop de iteração: %v", err)
	}

	return produtos, nil
}
