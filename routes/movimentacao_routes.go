package routes

import (
	"controle_estoque/handlers"
	"net/http"
)

func RegistrarRotasMovimentacao(mux *http.ServeMux, handler *handlers.MovimentacaoHandler) {
	mux.HandleFunc("GET /movimentacoes", handler.Listar)
	mux.HandleFunc("GET /produtos/{id}/movimentacoes", handler.ListarPorProduto)

	mux.HandleFunc("POST /produtos/{id}/entrada", handler.Entrada)
	mux.HandleFunc("POST /produtos/{id}/saida", handler.Saida)
}
