-- =====================================================
-- BANCO DE DADOS: controle_estoque
-- DESCRIÇÃO: Estrutura das tabelas do sistema de
-- controle de estoque.
-- =====================================================

-- TABELA DE PRODUTOS

CREATE TABLE produtos (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    codigo VARCHAR(50) NOT NULL UNIQUE,
    quantidade INTEGER NOT NULL DEFAULT 0,
    preco NUMERIC(10,2) NOT NULL,
    quantidade_minima INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_quantidade
        CHECK (quantidade >= 0),

    CONSTRAINT chk_preco
        CHECK (preco >= 0),

    CONSTRAINT chk_quantidade_minima
        CHECK (quantidade_minima >= 0)
);


-- TABELA DE MOVIMENTAÇÕES

CREATE TABLE movimentacoes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    produto_id INTEGER NOT NULL,
    tipo VARCHAR(10) NOT NULL,
    quantidade INTEGER NOT NULL,
    motivo VARCHAR(255),
    responsavel VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_produto
        FOREIGN KEY (produto_id)
        REFERENCES produtos(id),

    CONSTRAINT chk_tipo
        CHECK (tipo IN ('entrada', 'saida')),

    CONSTRAINT chk_quantidade_movimentacao
        CHECK (quantidade > 0)
);


-- Índice para facilitar consultas de movimentações

CREATE INDEX idx_movimentacoes_produto
ON movimentacoes(produto_id);


-- TABELA DE CLIENTES DA API

CREATE TABLE api_clients (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    key_hash CHAR(64) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);