First Go Project

# Primeiro Projeto Golang - API REST

[![Go](https://img.shields.io/badge/Go-1.25-blue)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-3.41.2-orange)](https://www.sqlite.org/index.html)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

Este foi meu **primeiro projeto em Golang**, construindo uma **API REST** simples para gerenciamento de tarefas (*tasks*).

---

## Tecnologias utilizadas

- **Golang**
- **Gorilla Mux** (para roteamento)
- **SQLite** (banco de dados)
- **Insomnia** (para testes de endpoints)

---

## Endpoints da API

### Criar uma nova task
**POST** `localhost:3030/tasks`  

**JSON de requisição:**
```json
{
    "title": "Golang initialize",
    "description": "First Task Golanf",
    "status": false
}
````

**JSON de requisição:**
```json
{
    "message": "Task criada com sucesso!"
}
````


### Listar todas as tasks
**GET** `localhost:3030/tasks`
````json
[
    {
        "id": 1,
        "title": "Golang initialize",
        "description": "First Task Golanf",
        "status": false
    }
]
````
### Atualizar uma task
**PUT** `localhost:3030/task/{id}`
````json
{
    "title": "Golang initializess",
    "description": "First Task Golanf",
    "status": false
}
````
**RESPOSTA** 
````json
{
    "message": "Task atualizada com sucesso!"
}
````
### Deletar uma Task
**DELETE** `localhost:3030/task/{id}`

**RESPOSTA**
````json
{
    "message": "Task deletada com sucesso!"
}
````

### Teste
**Os endpoints foram testados utilizando Insomnia, garantindo que as operações de CRUD funcionassem corretamente.**

# Como executar o projeto

Clone este repositório:
```bash
git clone https://github.com/SEU_USUARIO/NOME_DO_REPO.git

cd NOME_DO_REPO

go get -u github.com/gorilla/mux
go get -u modernc.org/sqlite
go run main.go
```
### Abra o Insomnia ou qualquer cliente REST e teste os endpoints em:
**Localhost:3030.**




