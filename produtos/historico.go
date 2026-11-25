package produtos

import (
	"database/sql"
	"fmt"
)

func RegistrarHistorico(db *sql.DB, id int, tipo string, quantidade int, responsavel string) error {
	query := `
        INSERT INTO movimentacoes
        (produto_id, tipo, quantidade, responsavel)
        VALUES ($1, $2, $3, $4)
    `

	_, err := db.Exec(
		query,
		id,
		tipo,
		quantidade,
		responsavel,
	)

	if err != nil {
		return fmt.Errorf("erro ao registrar histórico: %w", err)
	}

	return nil
}

func ListarHistorico(db *sql.DB) ([]Movimentacao, error) {
	query := `
        SELECT m.id, m.produto_id, COALESCE(p.nome, 'Produto excluído'),
               m.tipo, m.quantidade, m.responsavel, m.created_at
        FROM movimentacoes m
        LEFT JOIN produtos p ON p.id = m.produto_id
        ORDER BY m.created_at DESC, m.id DESC
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar histórico: %w", err)
	}
	defer rows.Close()

	var lista []Movimentacao

	for rows.Next() {
		var m Movimentacao

		err := rows.Scan(
			&m.ID,
			&m.ProdutoID,
			&m.ProdutoNome,
			&m.Tipo,
			&m.Quantidade,
			&m.Responsavel,
			&m.Data,
		)

		if err != nil {
			return nil, fmt.Errorf("erro ao ler movimentação: %w", err)
		}

		lista = append(lista, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro na consulta: %w", err)
	}

	return lista, nil
}
