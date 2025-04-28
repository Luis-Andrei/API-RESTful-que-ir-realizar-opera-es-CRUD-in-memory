# API Bancária em Go

Uma API RESTful em Go que realiza operações CRUD para clientes bancários (pessoais e corporativos) com armazenamento em memória.

## Funcionalidades

- Criação de clientes pessoais e corporativos
- Consulta de clientes por ID
- Listagem de todos os clientes
- Realização de saques
- Consulta de extrato

## Estrutura do Projeto

```
.
├── database/         # Implementações do banco de dados
├── handlers/         # Manipuladores HTTP
├── models/          # Modelos de dados
└── main.go          # Ponto de entrada da aplicação
```

## Requisitos

- Go 1.16 ou superior
- PostgreSQL (para testes)

## Instalação

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/api-bancaria-go.git
cd api-bancaria-go
```

2. Instale as dependências:
```bash
go mod download
```

3. Execute os testes:
```bash
go test ./...
```

## Configuração do Banco de Dados para Testes

1. Instale o PostgreSQL
2. Crie um banco de dados de teste:
```sql
CREATE DATABASE bank_test;
```
3. Configure as variáveis de ambiente:
```bash
export TEST_DB_HOST=localhost
export TEST_DB_PORT=5432
export TEST_DB_USER=postgres
export TEST_DB_PASSWORD=sua_senha
export TEST_DB_NAME=bank_test
```

## Executando a Aplicação

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`

## Endpoints

- `POST /api/clients/personal` - Cria um cliente pessoal
- `POST /api/clients/corporate` - Cria um cliente corporativo
- `GET /api/clients/:id` - Obtém um cliente por ID
- `GET /api/clients` - Lista todos os clientes
- `POST /api/clients/:id/withdraw` - Realiza um saque
- `GET /api/clients/:id/statement` - Obtém o extrato do cliente

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes. 