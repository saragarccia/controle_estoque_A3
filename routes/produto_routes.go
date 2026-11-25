package routes

import (
	"controle_estoque/handlers"
	"controle_estoque/middleware"
	"database/sql"
	"net/http"
)

func RegistrarRotas(mux *http.ServeMux, handler *handlers.ProdutoHandler, db *sql.DB) {
	mux.HandleFunc("GET /produtos", handler.Listar)
	mux.HandleFunc("GET /produtos/{id}", handler.Consultar)

	mux.Handle("POST /produtos", middleware.APIKey(db, http.HandlerFunc(handler.Cadastrar)))
	mux.Handle("PUT /produtos/{id}", middleware.APIKey(db, http.HandlerFunc(handler.Atualizar)))
	mux.Handle("DELETE /produtos/{id}", middleware.APIKey(db, http.HandlerFunc(handler.Excluir)))
}
