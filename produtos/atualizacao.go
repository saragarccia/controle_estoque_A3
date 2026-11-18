package produtos

import (
    "database/sql"
    "fmt"
)

func AtualizarProduto(db *sql.DB, p Produto) error {
    query := `
        UPDATE produtos
        SET nome = $1,
            descricao = $2,
            codigo = $3,
            preco = $4,
            quantidade_minima = $5
        WHERE id = $6
    `

    resultado, err := db.Exec(
        query,
        p.Nome,
        p.Descricao,
        p.Codigo,
        p.Preco,
        p.QuantidadeMinima,
        p.ID,
    )

    if err != nil {
        return fmt.Errorf("erro ao atualizar produto: %w", err)
    }

    linhas, err := resultado.RowsAffected()
    if err != nil {
        return fmt.Errorf("erro ao verificar atualização: %w", err)
    }

    if linhas == 0 {
        return fmt.Errorf("produto com ID %d não encontrado", p.ID)
    }

    fmt.Println("Produto atualizado com sucesso!")
    return nil
}