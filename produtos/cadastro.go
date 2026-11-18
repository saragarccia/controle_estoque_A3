package produtos

import (
    "database/sql"
    "fmt"
)

func CadastrarProduto(db *sql.DB, p Produto) error {
    query := `
        INSERT INTO produtos
        (nome, descricao, codigo, quantidade, preco, quantidade_minima)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `

    err := db.QueryRow(
        query,
        p.Nome,
        p.Descricao,
        p.Codigo,
        p.Quantidade,
        p.Preco,
        p.QuantidadeMinima,
    ).Scan(&p.ID)

    if err != nil {
        return fmt.Errorf("erro ao cadastrar produto: %w", err)
    }

    fmt.Println("Produto cadastrado com sucesso! ID:", p.ID)
    return nil
}