package produtos

import (
	"database/sql"
	"fmt"
)

func CadastrarProduto(db *sql.DB, p Produto) (int, error) {
	query := `
        INSERT INTO produtos
        (nome, descricao, codigo, quantidade, preco, quantidade_minima)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	var id int

	err := db.QueryRow(
		query,
		p.Nome,
		p.Descricao,
		p.Codigo,
		p.Quantidade,
		p.Preco,
		p.QuantidadeMinima,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao cadastrar produto: %w", err)
	}

	return id, nil
}
