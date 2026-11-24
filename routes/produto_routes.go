package routes

import (
	"controle_estoque/handlers"
	"net/http"
)

func RegistrarRotas(mux *http.ServeMux, handler *handlers.ProdutoHandler) {
	mux.HandleFunc("GET /produtos", handler.Listar)
	mux.HandleFunc("POST /produtos", handler.Cadastrar)

	mux.HandleFunc("GET /produtos/{id}", handler.Consultar)
	mux.HandleFunc("PUT /produtos/{id}", handler.Atualizar)
	mux.HandleFunc("DELETE /produtos/{id}", handler.Excluir)
}
