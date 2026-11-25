package repository

import (
	"controle_estoque/produtos"
	"database/sql"
)

type MovimentacaoRepository interface {
	Entrada(id int, quantidade int, responsavel string) error
	Saida(id int, quantidade int, responsavel string) error
	Listar() ([]produtos.Movimentacao, error)
	ListarPorProduto(id int) ([]produtos.Movimentacao, error)
}

type MovimentacaoRepositoryDB struct {
	DB *sql.DB
}

func (r *MovimentacaoRepositoryDB) Entrada(id int, quantidade int, responsavel string) error {
	return produtos.EntradaProduto(r.DB, id, quantidade, responsavel)
}

func (r *MovimentacaoRepositoryDB) Saida(id int, quantidade int, responsavel string) error {
	return produtos.SaidaProduto(r.DB, id, quantidade, responsavel)
}

func (r *MovimentacaoRepositoryDB) Listar() ([]produtos.Movimentacao, error) {
	return produtos.ListarHistorico(r.DB)
}

func (r *MovimentacaoRepositoryDB) ListarPorProduto(id int) ([]produtos.Movimentacao, error) {
	return produtos.ListarHistoricoProduto(r.DB, id)
}
