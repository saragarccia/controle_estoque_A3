package produtos

import (
	"database/sql"
	"fmt"
)

func ExcluirProduto(db *sql.DB, id int) error {
	if id <= 0 {
		return fmt.Errorf("ID do produto inválido")
	}

	query := `
        DELETE FROM produtos
        WHERE id = $1
    `

	resultado, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao excluir produto: %w", err)
	}

	linhas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar exclusão: %w", err)
	}

	if linhas == 0 {
		return fmt.Errorf("%w: ID %d", ErrProdutoNaoEncontrado, id)
	}

	fmt.Println("Produto excluído com sucesso!")
	return nil
}