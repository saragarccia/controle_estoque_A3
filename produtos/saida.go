package produtos

import (
	"database/sql"
	"fmt"
)

func SaidaProduto(db *sql.DB, id int, quantidade int, responsavel string) error {
	query := `
        UPDATE produtos
        SET quantidade = quantidade - $1
        WHERE id = $2 AND quantidade >= $1
    `

	resultado, err := db.Exec(
		query,
		quantidade,
		id,
	)

	if err != nil {
		return fmt.Errorf("erro ao remover produto: %w", err)
	}

	linhas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar remoção: %w", err)
	}

	if linhas == 0 {
		if _, err := ConsultarProduto(db, id); err != nil {
			return err
		}
		return ErrEstoqueInsuficiente
	}

	if err := RegistrarHistorico(db, id, "saida", quantidade, responsavel); err != nil {
		return err
	}

	fmt.Println("Produto removido com sucesso!")
	return nil
}
