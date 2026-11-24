package handlers

import (
	"controle_estoque/produtos"
	"controle_estoque/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ProdutoHandler struct {
	Repository repository.ProdutoRepository
}

func (h *ProdutoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	produtos, err := h.Repository.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar produtos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func (h *ProdutoHandler) Consultar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	produto, err := h.Repository.Consultar(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func (h *ProdutoHandler) Cadastrar(w http.ResponseWriter, r *http.Request) {
	var produto produtos.Produto

	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if produto.Nome == "" || produto.Codigo == "" ||
		produto.Quantidade < 0 || produto.Preco < 0 ||
		produto.QuantidadeMinima < 0 {
		http.Error(w, "Dados do produto inválidos", http.StatusBadRequest)
		return
	}

	id, err := h.Repository.Cadastrar(produto)
	if err != nil {
		http.Error(w, "Erro ao cadastrar produto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/produtos/%d", id))
	w.WriteHeader(http.StatusCreated)
}

func (h *ProdutoHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var produto produtos.Produto

	err = json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if produto.Nome == "" || produto.Codigo == "" ||
		produto.Preco < 0 || produto.QuantidadeMinima < 0 {
		http.Error(w, "Dados do produto inválidos", http.StatusBadRequest)
		return
	}

	produto.ID = id

	err = h.Repository.Atualizar(produto)
	if err != nil {
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProdutoHandler) Excluir(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.Repository.Excluir(id)
	if err != nil {
		http.Error(w, "Erro ao excluir produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
