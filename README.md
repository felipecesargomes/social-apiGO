# Social API GO

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)

<p align="center">
  <img src="https://github.com/user-attachments/assets/f27bab54-cecd-4c80-a268-1fda60127b79" alt="Go Logo" />
</p>


## Descrição

Projeto API Cadastro de usuários desenvolvido na linguagem Go, MySQL como banco de dados.

## Configuração (como executar com docker-composer)

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```
DB_TYPE=mysql
DB_HOST=db
DB_USER=root
DB_PASS=root
DB_NAME=devbook
DB_PORT=3306
API_PORT=":5000"
```

### Instalação

1 - Clone o repositório:
```
git clone https://github.com/seu-usuario/social-apiGO.git
```

2 - Navegue até o diretório do projeto:
```
cd social-apiGO
```

3 - Execute o docker-compose
```
docker-compose up --build
```

A API estará disponível em http://localhost:5000

### Endpoints

Usuários:
* POST - /users: Cria um novo usuário.
* GET - /users: Lista todos os usuários.
* GET - /users/{id}: Obtém um usuário pelo ID.
* PUT - /users/{id}: Atualiza um usuário pelo ID.
* DELETE - /users/{id}: Deleta um usuário pelo ID.
