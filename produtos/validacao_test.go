package produtos

import "testing"

func TestValidarProduto(t *testing.T) {
    p := Produto{
        Nome: "Teclado",
        Codigo: "TEC001",
        Preco: 50,
        QuantidadeMinima: 5,
    }

    if err := ValidarProduto(p); err != nil {
        t.Fatal("Produto válido foi rejeitado:", err)
    }

    p.Nome = ""

    if ValidarProduto(p) == nil {
        t.Fatal("Aceitou um produto sem nome")
    }

    p.Nome = "Teclado"
    p.Preco = -10

    if ValidarProduto(p) == nil {
        t.Fatal("Aceitou um preço negativo")
    }
}