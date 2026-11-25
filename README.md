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

É necessário ter o Go e o PostgreSQL instalados. Os comandos SQL abaixo devem ser executados no `psql` com o superusuário `postgres`.

### 6.1. Criar o banco de dados

```sql
CREATE DATABASE controle_estoque;
```

### 6.2. Aplicar o schema

Na pasta do projeto, crie as tabelas a partir do arquivo `schema.sql`:


```bash
psql -U postgres -d controle_estoque -f schema.sql
```

### 6.3. Criar o usuário da aplicação

Crie um usuário próprio:

```sql
CREATE USER app_estoque WITH PASSWORD 'sua_senha';
```

### 6.4. Conceder as permissões

O usuário da aplicação recebe apenas as permissões necessárias em cada tabela:

```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON produtos TO app_estoque;
GRANT SELECT, INSERT ON api_clients TO app_estoque;
```

- `produtos`: leitura, cadastro, atualização e exclusão, usadas pelas rotas da API.
- `api_clients`: cadastro de chaves, pelo comando `gerar-chave`, e consulta, pelo middleware de autenticação.

### 6.5. Configurar as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto, conforme o exemplo da seção 4, usando o usuário e a senha criados no passo 6.3.

### 6.6. Gerar uma chave de API

```bash
go run ./cmd/gerar-chave NOME_DO_CLIENTE
```

A chave é exibida uma única vez no terminal. Guarde-a: o banco armazena apenas o hash SHA-256 da chave.

### 6.7. Iniciar o servidor

```bash
go build ./...
```

Para executar:

```bash
go run main.go
```

O servidor fica disponível em `http://localhost:8080`.

## 7. Uso da API com curl

Os exemplos abaixo usam o `curl` com a sintaxe do Bash (Linux, macOS ou Git Bash no Windows). A opção `-i` mostra o código de status e os cabeçalhos da resposta.

Antes de executar os comandos, guarde a chave gerada no passo 6.6 em uma variável:


| Operação | Comando | Sucesso | Erros possíveis |
|---|---|---|---|
| Listar produtos | `curl -i http://localhost:8080/produtos` | `200 OK` com a lista em JSON | `500` |
| Consultar um produto | `curl -i http://localhost:8080/produtos/1` | `200 OK` com o produto em JSON | `400` ID inválido · `404` produto não encontrado · `500` |
| Cadastrar um produto | `curl -i -X POST http://localhost:8080/produtos -H "X-API-Key: $CHAVE" -H "Content-Type: application/json" -d '{"nome":"Teclado","descricao":"Teclado ABNT2","codigo":"TEC01","quantidade":10,"preco":199.90,"quantidade_minima":2}'` | `201 Created` com o cabeçalho `Location: /produtos/{id}` | `400` dados inválidos · `401` chave ausente ou inválida · `500` |
| Atualizar um produto | `curl -i -X PUT http://localhost:8080/produtos/1 -H "X-API-Key: $CHAVE" -H "Content-Type: application/json" -d '{"nome":"Teclado sem fio","descricao":"Teclado ABNT2 sem fio","codigo":"TEC01","preco":249.90,"quantidade_minima":3}'` | `204 No Content` | `400` ID ou dados inválidos · `401` chave ausente ou inválida · `404` produto não encontrado · `500` |
| Excluir um produto | `curl -i -X DELETE http://localhost:8080/produtos/1 -H "X-API-Key: $CHAVE"` | `204 No Content` | `400` ID inválido · `401` chave ausente ou inválida · `404` produto não encontrado · `500` |

Observações:

- Troque o `1` das URLs pelo ID do produto desejado. O ID de um produto recém-cadastrado aparece no cabeçalho `Location` da resposta.
- No cadastro, `nome` e `codigo` são obrigatórios, e `quantidade`, `preco` e `quantidade_minima` não podem ser negativos.
- A atualização não altera a `quantidade` em estoque. Ela só muda por entradas e saídas de produtos.
- O `codigo` não pode se repetir entre produtos.

No PowerShell do Windows, use `curl.exe` no lugar de `curl`, defina a chave com `$CHAVE = "cole_aqui_a_chave_gerada"` e escape as aspas internas do JSON com barra invertida, por exemplo `-d '{\"nome\":\"Teclado\", ...}'`.