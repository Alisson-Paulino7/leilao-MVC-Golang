package models


import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UpProduct struct {
	Id int 			`json:"id"`
	Novolance int 	`json:"novolance"`
}

func (u *UpProduct) NovoLance() error {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	result, err := db.Exec("UPDATE produtos SET lance_atual = ? WHERE id = ?",
		u.Novolance, u.Id)

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