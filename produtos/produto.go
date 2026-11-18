package produtos

type Produto struct {
    ID               int
    Nome             string
    Descricao        string
    Codigo           string
    Quantidade       int
    Preco            float64
    QuantidadeMinima int
}