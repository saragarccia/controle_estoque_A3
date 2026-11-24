package produtos

type Produto struct {
	ID               int     `json:"id"`
	Nome             string  `json:"nome"`
	Descricao        string  `json:"descricao"`
	Codigo           string  `json:"codigo"`
	Quantidade       int     `json:"quantidade"`
	Preco            float64 `json:"preco"`
	QuantidadeMinima int     `json:"quantidade_minima"`
}