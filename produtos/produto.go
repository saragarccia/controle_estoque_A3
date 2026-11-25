package produtos

import "time"

type Produto struct {
	ID               int     `json:"id"`
	Nome             string  `json:"nome"`
	Descricao        string  `json:"descricao"`
	Codigo           string  `json:"codigo"`
	Quantidade       int     `json:"quantidade"`
	Preco            float64 `json:"preco"`
	QuantidadeMinima int     `json:"quantidade_minima"`
}

type Movimentacao struct {
	ID          int       `json:"id"`
	ProdutoID   int       `json:"produto_id"`
	ProdutoNome string    `json:"produto_nome"`
	Tipo        string    `json:"tipo"`
	Quantidade  int       `json:"quantidade"`
	Responsavel string    `json:"responsavel"`
	Data        time.Time `json:"data"`
}
