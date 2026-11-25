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

A conexão com o PostgreSQL é feita pelo arquivo `config/database.go`, com o driver `pgx`. Os dados de acesso não ficam no código: eles são lidos da variável de ambiente `DATABASE_URL`, no formato:

```
postgres://USUARIO:SENHA@HOST:PORTA/BANCO?sslmode=disable
```

Exemplo com o usuário criado no passo 6.3:

```
postgres://app_estoque:sua_senha@localhost:5432/controle_estoque?sslmode=disable
```

Caracteres especiais na senha (como `@`, `#`, `/` ou `:`) devem ser codificados na URL. Por exemplo, `@` vira `%40` e `#` vira `%23`.
## 5. Estrutura do projeto

```text
controle_estoque_A3/
├── cmd/
│   └── gerar-chave/     # Comando para gerar chaves de API
├── config/              # Conexão com o banco (DATABASE_URL)
├── handlers/            # Handlers HTTP (produtos, movimentações e health)
├── middleware/          # Autenticação por API Key
├── produtos/            # Modelos, validação e acesso ao banco (SQL)
├── repository/          # Interfaces e implementações dos DAOs
├── routes/              # Registro das rotas
├── .gitignore
├── go.mod
├── go.sum
├── main.go              # Lê a configuração e monta o sistema
├── README.md
└── schema.sql           # Estrutura do banco de dados
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

```bash
psql -U postgres
```
```sql
CREATE USER app_estoque WITH PASSWORD 'sua_senha';
```

### 6.4. Conceder as permissões

As permissões são concedidas nas tabelas do banco `controle_estoque`, por isso é necessário estar conectado a ele. Ainda no `psql`, troque de banco:

```sql
\c controle_estoque
```

O prompt deve mudar para `controle_estoque=#`. Em seguida, conceda ao usuário da aplicação apenas as permissões necessárias em cada tabela:

```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON produtos TO app_estoque;
GRANT SELECT, INSERT ON movimentacoes TO app_estoque;
GRANT SELECT, INSERT ON api_clients TO app_estoque;
```

- `produtos`: leitura, cadastro, atualização e exclusão, usadas pelas rotas da API.
- `movimentacoes`: registro das entradas e saídas e consulta do histórico.
- `api_clients`: cadastro de chaves, pelo comando `gerar-chave`, e consulta, pelo middleware de autenticação.

Para conferir as permissões concedidas:

```sql
SELECT table_name, privilege_type FROM information_schema.role_table_grants WHERE grantee = 'app_estoque' ORDER BY table_name;
```

O resultado deve listar as três tabelas com as permissões acima. Para sair do `psql`, digite `\q`.

> Se aparecer `relação "produtos" não existe`, o `psql` não está conectado ao banco `controle_estoque`: confira se o prompt mostra `controle_estoque=#` e se o passo 6.2 foi executado.

### 6.5. Configurar as variáveis de ambiente

Defina a variável `DATABASE_URL` com o usuário e a senha criados no passo 6.3. Ela vale apenas para o terminal em que foi definida, por isso execute os passos 6.6 e 6.7 nesse mesmo terminal.

No Linux, macOS ou Git Bash:

```bash
export DATABASE_URL="postgres://app_estoque:sua_senha@localhost:5432/controle_estoque?sslmode=disable"
```

No PowerShell (Windows):

```powershell
$env:DATABASE_URL = "postgres://app_estoque:sua_senha@localhost:5432/controle_estoque?sslmode=disable"
```

Se a variável não estiver definida, o servidor encerra com a mensagem `variável de ambiente DATABASE_URL não definida`.
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
| Registrar entrada no estoque | `curl -i -X POST http://localhost:8080/produtos/1/entrada -H "X-API-Key: $CHAVE" -H "Content-Type: application/json" -d '{"quantidade":10,"responsavel":"Rafa"}'` | `204 No Content` | `400` ID inválido ou quantidade menor ou igual a zero · `401` chave ausente ou inválida · `404` produto não encontrado · `500` |
| Registrar saída do estoque | `curl -i -X POST http://localhost:8080/produtos/1/saida -H "X-API-Key: $CHAVE" -H "Content-Type: application/json" -d '{"quantidade":3,"responsavel":"Rafa"}'` | `204 No Content` | `400` ID inválido ou quantidade menor ou igual a zero · `401` chave ausente ou inválida · `404` produto não encontrado · `409` estoque insuficiente · `500` |
| Histórico de um produto | `curl -i http://localhost:8080/produtos/1/movimentacoes` | `200 OK` com as movimentações do produto em JSON, da mais recente para a mais antiga | `400` ID inválido · `404` produto não encontrado · `500` |
| Histórico completo | `curl -i http://localhost:8080/movimentacoes` | `200 OK` com todas as movimentações em JSON, da mais recente para a mais antiga | `500` |
Observações:

- Troque o `1` das URLs pelo ID do produto desejado. O ID de um produto recém-cadastrado aparece no cabeçalho `Location` da resposta.
- No cadastro, `nome` e `codigo` são obrigatórios, e `quantidade`, `preco` e `quantidade_minima` não podem ser negativos.
- A atualização não altera a `quantidade` em estoque. Ela só muda por entradas e saídas de produtos.
- O `codigo` não pode se repetir entre produtos.

No PowerShell do Windows, use `curl.exe` no lugar de `curl`, defina a chave com `$CHAVE = "cole_aqui_a_chave_gerada"` e escape as aspas internas do JSON com barra invertida, por exemplo `-d '{\"nome\":\"Teclado\", ...}'`.