package produtos

import (
    "database/sql"
    "fmt"
)

func ListarProdutos(db *sql.DB) ([]Produto, error) {
    query := `
        SELECT id, nome, descricao, codigo,
               quantidade, preco, quantidade_minima
        FROM produtos
        ORDER BY id
    `

    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("erro ao listar produtos: %w", err)
    }
    defer rows.Close()

    var lista []Produto

    for rows.Next() {
        var p Produto
        var descricao sql.NullString

        err := rows.Scan(
            &p.ID,
            &p.Nome,
            &descricao,
            &p.Codigo,
            &p.Quantidade,
            &p.Preco,
            &p.QuantidadeMinima,
        )

        if err != nil {
            return nil, fmt.Errorf("erro ao ler produto: %w", err)
        }

        if descricao.Valid {
            p.Descricao = descricao.String
        }

        lista = append(lista, p)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("erro na consulta: %w", err)
    }

    return lista, nil
}