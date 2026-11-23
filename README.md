# Sistema de Controle de Estoque

Sistema de controle de estoque desenvolvido em Go, com integração ao banco de dados PostgreSQL.

## 1. Tecnologias utilizadas

- **Go:** linguagem utilizada no desenvolvimento do sistema.
- **PostgreSQL:** sistema gerenciador do banco de dados.
- **Git e GitHub:** ferramentas utilizadas para versionamento e compartilhamento do projeto.

## 2. Banco de dados

O projeto utiliza o banco de dados `controle_estoque`, responsável por armazenar as informações dos produtos e das movimentações de estoque.

O banco possui três tabelas:

### 2.1. produtos

Armazena as informações dos produtos cadastrados.

- `id`: identificador automático do produto.
- `nome`: nome do produto.
- `descricao`: descrição do produto.
- `codigo`: código único do produto.
- `quantidade`: quantidade disponível em estoque.
- `preco`: preço do produto.
- `quantidade_minima`: quantidade mínima definida para o estoque.
- `created_at`: data e hora do cadastro.

### 2.2. movimentacoes

Armazena os registros de movimentações dos produtos.

- `id`: identificador da movimentação.
- `produto_id`: identifica o produto relacionado.
- `tipo`: indica se a movimentação é uma entrada ou saída.
- `quantidade`: quantidade movimentada.
- `motivo`: motivo da movimentação.
- `created_at`: data e hora do registro.

### 2.3. api_clients

Tabela preparada para armazenar os clientes utilizados na autenticação da API.

- `id`: identificador do cliente.
- `name`: nome do cliente.
- `key_hash`: hash da chave de API.
- `created_at`: data e hora do cadastro.

## 3. Regras do banco de dados

O banco possui regras para manter a integridade das informações:

- Os identificadores são gerados automaticamente.
- O nome e o código do produto são obrigatórios.
- O código do produto não pode ser repetido.
- A quantidade, o preço e a quantidade mínima não podem ser negativos.
- As movimentações precisam estar relacionadas a um produto existente.
- O tipo de movimentação deve ser entrada ou saída.
- A quantidade de uma movimentação deve ser maior que zero.

## 4. Configuração da conexão

A conexão com o PostgreSQL é realizada por meio do arquivo `config/database.go`.

As configurações de acesso são armazenadas em um arquivo `.env` local, que não deve ser enviado ao GitHub.

Exemplo das variáveis necessárias:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=app_estoque
DB_PASSWORD=sua_senha
DB_NAME=controle_estoque
DB_SSLMODE=disable
```

## 5. Estrutura do projeto

```text
controle_estoque_A3/
├── config/
├── produtos/
├── .env                 # Arquivo local, não versionado
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── README.md
└── schema.sql
```

O arquivo `schema.sql` contém os comandos SQL utilizados para criar a estrutura do banco de dados.

## 6. Execução do projeto

É necessário ter o Go e o PostgreSQL instalados e configurar corretamente o arquivo `.env`.

Para verificar a compilação:

```bash
go build ./...
```

Para executar:

```bash
go run main.go
```