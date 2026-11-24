package produtos

import (
	"database/sql"
	"fmt"
)

func SaidaProduto(db *sql.DB, id int, quantidade int) error {
	query := `
        UPDATE produtos
        SET quantidade = quantidade - $1
        WHERE id = $2
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
		return fmt.Errorf("produto com ID %d não encontrado", id)
	}

	fmt.Println("Produto removido com sucesso!")
	return nil
}
