package repository

import (
	"controle_estoque/produtos"
	"database/sql"
)

type ProdutoRepository interface {
	Cadastrar(p produtos.Produto) (int, error)
	Consultar(id int) (produtos.Produto, error)
	Listar() ([]produtos.Produto, error)
	Atualizar(p produtos.Produto) error
	Excluir(id int) error
}

type ProdutoRepositoryDB struct {
	DB *sql.DB
}

func (r *ProdutoRepositoryDB) Cadastrar(p produtos.Produto) (int, error) {
	produto, err := produtos.CadastrarProduto(r.DB, p)
	if err != nil {
		return 0, err
	}
	return produto.ID, nil
}

func (r *ProdutoRepositoryDB) Consultar(id int) (produtos.Produto, error) {
	return produtos.ConsultarProduto(r.DB, id)
}

func (r *ProdutoRepositoryDB) Listar() ([]produtos.Produto, error) {
	return produtos.ListarProdutos(r.DB)
}

func (r *ProdutoRepositoryDB) Atualizar(p produtos.Produto) error {
	return produtos.AtualizarProduto(r.DB, p)
}

func (r *ProdutoRepositoryDB) Excluir(id int) error {
	return produtos.ExcluirProduto(r.DB, id)
}