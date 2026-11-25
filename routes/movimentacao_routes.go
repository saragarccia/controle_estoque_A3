package routes

import (
	"controle_estoque/handlers"
	"controle_estoque/middleware"
	"database/sql"
	"net/http"
)

func RegistrarRotasMovimentacao(mux *http.ServeMux, handler *handlers.MovimentacaoHandler, db *sql.DB) {
	mux.HandleFunc("GET /movimentacoes", handler.Listar)
	mux.HandleFunc("GET /produtos/{id}/movimentacoes", handler.ListarPorProduto)

	mux.Handle("POST /produtos/{id}/entrada", middleware.APIKey(db, http.HandlerFunc(handler.Entrada)))
	mux.Handle("POST /produtos/{id}/saida", middleware.APIKey(db, http.HandlerFunc(handler.Saida)))
}
