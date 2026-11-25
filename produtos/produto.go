package produtos

import "time"

type Produto struct {
	ID               int     json:"id"
	Nome             string  json:"nome"
	Descricao        string  json:"descricao"
	Codigo           string  json:"codigo"
	Quantidade       int     json:"quantidade"
	Preco            float64 json:"preco"
	QuantidadeMinima int     json:"quantidade_minima"
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