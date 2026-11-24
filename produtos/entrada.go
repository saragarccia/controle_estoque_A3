package produtos

import (
	"database/sql"
	"fmt"
)

func EntradaProduto(db *sql.DB, id int, quantidade int) error {
	query := `
        UPDATE produtos
        SET quantidade = quantidade + $1
        WHERE id = $2
    `

	resultado, err := db.Exec(
		query,
		quantidade,
		id,
	)

	if err != nil {
		return fmt.Errorf("erro ao abastecer produto: %w", err)
	}

	linhas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar abastecimento: %w", err)
	}

	if linhas == 0 {
		return fmt.Errorf("produto com ID %d não encontrado", id)
	}

	fmt.Println("Produto abastecido com sucesso!")
	return nil
}
