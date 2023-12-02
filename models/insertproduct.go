package models


import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type InsProduto struct {
	Nome         string 	`json:"nome"`
	Descricao	 string 	`json:"descricao"`
	LanceInicial int 		`json:"lance_inicial"`
	TmpExpiracao int 		`json:"tmp_expiracao"`
	// DataCadastro string   	`json:"data_cadastro"`
	// DataExpires string 		`json:"data_expires"`
	FotoProduct  []byte 	`json:"Foto_Product"`
}



func ( i *InsProduto) SalvarProduto() error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	result, err := db.Exec("insert into produtos(nome_product, desc_product, lance_inicial, tmp_expiracao, foto_product) VALUES (?, ?, ?, ?, ?)",
		i. Nome, i.Descricao, i.LanceInicial, i.TmpExpiracao,  i.FotoProduct)

	if err != nil {
		return fmt.Errorf("erro ao inserir foto no banco de dados: %v", err)
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
