# Social API GO

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)

## Descrição

Este projeto é uma API social desenvolvida em Go, utilizando o framework Gorilla Mux para roteamento e MySQL como banco de dados. A API permite a criação, leitura, atualização e exclusão de usuários, além de autenticação de login.

## Estrutura do Projeto

TODO: Criar árvore do projeto.

## Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

DB_TYPE=mysql
DB_HOST=127.0.0.1
DB_USER=root
DB_PASS=root
DB_NAME=devbook
DB_PORT=3306
API_PORT=":5000"

### Banco de Dados

Execute o script SQL localizado em `sql/sql.sql` para criar o banco de dados e a tabela de usuários:

```sql
CREATE DATABASE IF NOT EXISTS DEVBOOK;

USE DEVBOOK;

DROP TABLE IF EXISTS USERS;

CREATE TABLE USERS (
    id INT PRIMARY KEY AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL unique,
    email VARCHAR(50) NOT NULL unique,
    password VARCHAR(50) NOT NULL,
    CREATED_AT DATETIME default CURRENT_TIMESTAMP()
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

### Instalação

1 - Clone o repositório:
git clone https://github.com/seu-usuario/social-apiGO.git

2 - Navegue até o diretório do projeto:
cd social-apiGO

3 - Instale as dependências:
go mod tidy

### Executando a Aplicação

Para rodar a aplicação, utilize o comando:

go run main.go

A API estará disponível em http://localhost:5000

Endpoints
Autenticação
POST /login: Autentica um usuário.
Usuários
POST /users: Cria um novo usuário.
GET /users: Lista todos os usuários.
GET /users/{id}: Obtém um usuário pelo ID.
PUT /users/{id}: Atualiza um usuário pelo ID.
DELETE /users/{id}: Deleta um usuário pelo ID.