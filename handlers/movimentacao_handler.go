package handlers

import (
	"controle_estoque/produtos"
	"controle_estoque/repository"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

type MovimentacaoHandler struct {
	Repository repository.MovimentacaoRepository
}

type dadosMovimentacao struct {
	Quantidade  int    `json:"quantidade"`
	Responsavel string `json:"responsavel"`
}

func (h *MovimentacaoHandler) Entrada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	dados, ok := lerDadosMovimentacao(w, r)
	if !ok {
		return
	}

	err = h.Repository.Entrada(id, dados.Quantidade, dados.Responsavel)

	if errors.Is(err, produtos.ErrProdutoNaoEncontrado) {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro ao registrar entrada", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MovimentacaoHandler) Saida(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	dados, ok := lerDadosMovimentacao(w, r)
	if !ok {
		return
	}

	err = h.Repository.Saida(id, dados.Quantidade, dados.Responsavel)

	if errors.Is(err, produtos.ErrProdutoNaoEncontrado) {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	if errors.Is(err, produtos.ErrEstoqueInsuficiente) {
		http.Error(w, "Estoque insuficiente para a saída", http.StatusConflict)
		return
	}

	if err != nil {
		http.Error(w, "Erro ao registrar saída", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MovimentacaoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	movimentacoes, err := h.Repository.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar movimentações", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movimentacoes)
}

func (h *MovimentacaoHandler) ListarPorProduto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	movimentacoes, err := h.Repository.ListarPorProduto(id)

	if errors.Is(err, produtos.ErrProdutoNaoEncontrado) {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro ao listar movimentações", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movimentacoes)
}

// lerDadosMovimentacao lê e valida o corpo da entrada/saída; se for inválido, responde 400.
func lerDadosMovimentacao(w http.ResponseWriter, r *http.Request) (dadosMovimentacao, bool) {
	var dados dadosMovimentacao

	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return dados, false
	}

	if dados.Quantidade <= 0 {
		http.Error(w, "A quantidade deve ser maior que zero", http.StatusBadRequest)
		return dados, false
	}

	dados.Responsavel = strings.TrimSpace(dados.Responsavel)
	if dados.Responsavel == "" {
		dados.Responsavel = "Sem responsável"
	}

	if utf8.RuneCountInString(dados.Responsavel) > 100 {
		http.Error(w, "O responsável deve ter no máximo 100 caracteres", http.StatusBadRequest)
		return dados, false
	}

	return dados, true
}
