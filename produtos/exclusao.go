package produtos

import (
    "database/sql"
    "fmt"
)

func ExcluirProduto(db *sql.DB, id int) error {
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
        return fmt.Errorf("erro ao verificar exclusao: %w", err)
    }

    if linhas == 0 {
        return fmt.Errorf("produto com ID %d nao encontrado", id)
    }

    fmt.Println("Produto excluido com sucesso!")
    return nil
}