# API CRUD em Go

Uma API RESTful simples construída em Go que realiza operações CRUD com armazenamento em memória.

## Funcionalidades

- Operações de Criar, Ler, Atualizar e Deletar
- Armazenamento em memória
- Endpoints RESTful
- Formato de requisição/resposta em JSON

## Endpoints da API

- `POST /itens` - Criar um novo item
- `GET /itens` - Obter todos os itens
- `GET /itens/{id}` - Obter um item específico
- `PUT /itens/{id}` - Atualizar um item
- `DELETE /itens/{id}` - Deletar um item

## Estrutura do Item

```json
{
    "id": "string",
    "nome": "string",
    "preco": "float64"
}
```

## Como Começar

1. Instale o Go (versão 1.21 ou superior)
2. Clone este repositório
3. Execute `go mod tidy` para instalar as dependências
4. Execute `go run main.go` para iniciar o servidor
5. O servidor iniciará na porta 8080

## Exemplos de Uso

### Criar um Item
```bash
curl -X POST http://localhost:8080/itens \
-H "Content-Type: application/json" \
-d '{"id": "1", "nome": "Item Teste", "preco": 19.99}'
```

### Obter Todos os Itens
```bash
curl http://localhost:8080/itens
```

### Obter um Item Específico
```bash
curl http://localhost:8080/itens/1
```

### Atualizar um Item
```bash
curl -X PUT http://localhost:8080/itens/1 \
-H "Content-Type: application/json" \
-d '{"id": "1", "nome": "Item Atualizado", "preco": 29.99}'
```

### Deletar um Item
```bash
curl -X DELETE http://localhost:8080/itens/1
``` 