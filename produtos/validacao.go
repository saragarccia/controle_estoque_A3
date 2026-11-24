package produtos

import (
    "fmt"
    "strings"
    "math"
)

func ValidarProduto(p Produto) error {
    if strings.TrimSpace(p.Nome) == "" {
        return fmt.Errorf("o nome do produto é obrigatório")
    }

    if strings.TrimSpace(p.Codigo) == "" {
        return fmt.Errorf("o código do produto é obrigatório")
    }

    if p.Preco < 0 || math.IsNaN(p.Preco) || math.IsInf(p.Preco, 0) {
        return fmt.Errorf("preço inválido")
    }

    if p.QuantidadeMinima < 0 {
        return fmt.Errorf("a quantidade mínima não pode ser negativa")
    }

    return nil
}