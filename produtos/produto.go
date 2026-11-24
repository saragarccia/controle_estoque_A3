package produtos

import "time"

type Produto struct {
	ID               int
	Nome             string
	Descricao        string
	Codigo           string
	Quantidade       int
	Preco            float64
	QuantidadeMinima int
}

type Movimentacao struct {
	ID          int
	ProdutoID   int
	ProdutoNome string
	Tipo        string
	Quantidade  int
	Responsavel string
	Data        time.Time
}
