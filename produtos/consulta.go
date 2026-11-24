package produtos

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrProdutoNaoEncontrado = errors.New("produto não encontrado")

func ConsultarProduto(db *sql.DB, id int) (Produto, error) {
	var p Produto
	var descricao sql.NullString

	query := `
        SELECT id, nome, descricao, codigo,
               quantidade, preco, quantidade_minima
        FROM produtos
        WHERE id = $1
    `

	err := db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Nome,
		&descricao,
		&p.Codigo,
		&p.Quantidade,
		&p.Preco,
		&p.QuantidadeMinima,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return Produto{}, fmt.Errorf("%w: ID %d", ErrProdutoNaoEncontrado, id)
	}

	if err != nil {
		return Produto{}, fmt.Errorf("erro ao consultar produto: %w", err)
	}

	if descricao.Valid {
		p.Descricao = descricao.String
	}

	return p, nil
}
