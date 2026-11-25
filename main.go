package main

import (
	"controle_estoque/config"
	"controle_estoque/handlers"
	"controle_estoque/repository"
	"controle_estoque/routes"
	"fmt"
	"net/http"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco:", err)
		return
	}

	defer db.Close()

	fmt.Println("Conexão com o PostgreSQL realizada com sucesso!")

	repositorio := &repository.ProdutoRepositoryDB{DB: db}
	handler := &handlers.ProdutoHandler{Repository: repositorio}
	movimentacaoRepositorio := &repository.MovimentacaoRepositoryDB{DB: db}
	movimentacaoHandler := &handlers.MovimentacaoHandler{Repository: movimentacaoRepositorio}

	mux := http.NewServeMux()
	routes.RegistrarRotas(mux, handler, db)
	routes.RegistrarRotasMovimentacao(mux, movimentacaoHandler, db)

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
