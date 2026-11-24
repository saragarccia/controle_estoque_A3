package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"controle_estoque/config"
	"controle_estoque/produtos"
)

var leitor = bufio.NewReader(os.Stdin)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco:", err)
		return
	}
	defer db.Close()

	fmt.Println("Conexão com o PostgreSQL realizada com sucesso!")

	for {
		fmt.Println()
		fmt.Println(" CONTROLE DE ESTOQUE ")
		fmt.Println("1 - Cadastrar produto")
		fmt.Println("2 - Listar produtos")
		fmt.Println("3 - Consultar produto por ID")
		fmt.Println("4 - Atualizar produto")
		fmt.Println("5 - Excluir produto")
		fmt.Println("6 - Reabastecer produto")
		fmt.Println("7 - Dar saida em estoque")
		fmt.Println("0 - Sair")

		switch lerTexto("Escolha uma opção: ") {
		case "1":
			cadastrar(db)
		case "2":
			listar(db)
		case "3":
			consultar(db)
		case "4":
			atualizar(db)
		case "5":
			excluir(db)
		case "6":
			darEntrada(db)
		case "7":
			darSaida(db)
		case "0":
			fmt.Println("Encerrando...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func cadastrar(db *sql.DB) {
	p := produtos.Produto{
		Nome:             lerTexto("Nome: "),
		Descricao:        lerTexto("Descrição: "),
		Codigo:           lerTexto("Código: "),
		Quantidade:       lerInt("Quantidade: "),
		Preco:            lerFloat("Preço: "),
		QuantidadeMinima: lerInt("Quantidade mínima: "),
	}

	if err := produtos.CadastrarProduto(db, p); err != nil {
		fmt.Println(err)
	}
}

func listar(db *sql.DB) {
	lista, err := produtos.ListarProdutos(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(lista) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}

	for _, p := range lista {
		imprimirProduto(p)
	}
}

func consultar(db *sql.DB) {
	p, err := produtos.ConsultarProduto(db, lerInt("ID do produto: "))
	if err != nil {
		fmt.Println(err)
		return
	}
	imprimirProduto(p)
}

func atualizar(db *sql.DB) {
	id := lerInt("ID do produto: ")

	atual, err := produtos.ConsultarProduto(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deixe em branco para manter o valor atual.")
	atual.Nome = lerTextoPadrao("Nome", atual.Nome)
	atual.Descricao = lerTextoPadrao("Descrição", atual.Descricao)
	atual.Codigo = lerTextoPadrao("Código", atual.Codigo)
	atual.Preco = lerFloatPadrao("Preço", atual.Preco)
	atual.QuantidadeMinima = lerIntPadrao("Quantidade mínima", atual.QuantidadeMinima)

	if err := produtos.AtualizarProduto(db, atual); err != nil {
		fmt.Println(err)
	}
}

func excluir(db *sql.DB) {
	id := lerInt("ID do produto: ")

	if strings.ToLower(lerTexto("Confirma a exclusão? (s/n): ")) != "s" {
		fmt.Println("Exclusão cancelada.")
		return
	}

	if err := produtos.ExcluirProduto(db, id); err != nil {
		fmt.Println(err)
	}
}

func imprimirProduto(p produtos.Produto) {
	fmt.Printf("[%d] %s (código %s) | qtd: %d | mín: %d | R$ %.2f",
		p.ID, p.Nome, p.Codigo, p.Quantidade, p.QuantidadeMinima, p.Preco)
	if p.Quantidade < p.QuantidadeMinima {
		fmt.Print("  ⚠ abaixo do mínimo")
	}
	fmt.Println()
	if p.Descricao != "" {
		fmt.Println("    ", p.Descricao)
	}
}

func lerTexto(msg string) string {
	fmt.Print(msg)
	texto, _ := leitor.ReadString('\n')
	return strings.TrimSpace(texto)
}

func lerInt(msg string) int {
	for {
		n, err := strconv.Atoi(lerTexto(msg))
		if err == nil {
			return n
		}
		fmt.Println("Digite um número inteiro válido.")
	}
}

func lerFloat(msg string) float64 {
	for {
		texto := strings.Replace(lerTexto(msg), ",", ".", 1)
		f, err := strconv.ParseFloat(texto, 64)
		if err == nil {
			return f
		}
		fmt.Println("Digite um valor numérico válido (ex: 10.50).")
	}
}

func lerTextoPadrao(campo, atual string) string {
	texto := lerTexto(fmt.Sprintf("%s [%s]: ", campo, atual))
	if texto == "" {
		return atual
	}
	return texto
}

func lerIntPadrao(campo string, atual int) int {
	for {
		texto := lerTexto(fmt.Sprintf("%s [%d]: ", campo, atual))
		if texto == "" {
			return atual
		}
		if n, err := strconv.Atoi(texto); err == nil {
			return n
		}
		fmt.Println("Digite um número inteiro válido.")
	}
}

func lerFloatPadrao(campo string, atual float64) float64 {
	for {
		texto := lerTexto(fmt.Sprintf("%s [%.2f]: ", campo, atual))
		if texto == "" {
			return atual
		}
		if f, err := strconv.ParseFloat(strings.Replace(texto, ",", ".", 1), 64); err == nil {
			return f
		}
		fmt.Println("Digite um valor numérico válido (ex: 10.50).")
	}
}

func darEntrada(db *sql.DB) {
	lista, err := produtos.ListarProdutos(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(lista) == 0 {
		fmt.Println("Nenhum produto adicionado ao estoque.")
		return
	}

	for _, p := range lista {
		imprimirProduto(p)
	}

	id := lerInt("ID do produto a reabastecer: ")

	quantidade := lerInt("Quantidade de novos produtos: ")
	for quantidade <= 0 {
		fmt.Println("A quantidade deve ser maior que zero")
		quantidade = lerInt("Quantidade de novos produtos: ")
	}
	if err := produtos.EntradaProduto(db, id, quantidade); err != nil {
		fmt.Println(err)
	}
}
func darSaida(db *sql.DB) {
	lista, err := produtos.ListarProdutos(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(lista) == 0 {
		fmt.Println("Nenhum removido do estoque.")
		return
	}

	for _, p := range lista {
		imprimirProduto(p)
	}

	id := lerInt("ID do produto a remover: ")

	quantidade := lerInt("Quantidade de remoções: ")
	for quantidade <= 0 {
		fmt.Println("A quantidade deve ser maior que zero")
		quantidade = lerInt("Quantidade de remoções: ")
	}
	if err := produtos.SaidaProduto(db, id, quantidade); err != nil {
		fmt.Println(err)
	}
}
