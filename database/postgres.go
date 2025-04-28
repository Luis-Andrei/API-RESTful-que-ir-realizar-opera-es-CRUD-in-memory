package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Luis-Andrei/api-users/models"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(host, port, user, password, dbname string) (Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) InitTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS clients (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			balance DECIMAL(15,2) NOT NULL,
			client_type VARCHAR(10) NOT NULL,
			cpf VARCHAR(14),
			cnpj VARCHAR(18),
			transactions JSONB
		)`,
	}

	for _, query := range queries {
		_, err := p.db.Exec(query)
		if err != nil {
			return fmt.Errorf("error creating tables: %v", err)
		}
	}

	return nil
}

func (p *PostgresDB) CreatePersonalClient(client *models.PersonalClient) error {
	transactions, err := json.Marshal(client.Transactions)
	if err != nil {
		return fmt.Errorf("error marshaling transactions: %v", err)
	}

	query := `
		INSERT INTO clients (id, name, balance, client_type, cpf, transactions)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = p.db.Exec(query,
		client.ID,
		client.Name,
		client.Balance,
		"personal",
		client.CPF,
		transactions)

	if err != nil {
		return fmt.Errorf("error creating personal client: %v", err)
	}

	return nil
}

func (p *PostgresDB) CreateCorporateClient(client *models.CorporateClient) error {
	transactions, err := json.Marshal(client.Transactions)
	if err != nil {
		return fmt.Errorf("error marshaling transactions: %v", err)
	}

	query := `
		INSERT INTO clients (id, name, balance, client_type, cnpj, transactions)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = p.db.Exec(query,
		client.ID,
		client.Name,
		client.Balance,
		"corporate",
		client.CNPJ,
		transactions)

	if err != nil {
		return fmt.Errorf("error creating corporate client: %v", err)
	}

	return nil
}

func (p *PostgresDB) GetClient(id string) (models.Client, error) {
	query := `
		SELECT id, name, balance, client_type, cpf, cnpj, transactions
		FROM clients
		WHERE id = $1`

	var (
		clientID     string
		name         string
		balance      float64
		clientType   string
		cpf          sql.NullString
		cnpj         sql.NullString
		transactions []byte
	)

	err := p.db.QueryRow(query, id).Scan(
		&clientID,
		&name,
		&balance,
		&clientType,
		&cpf,
		&cnpj,
		&transactions)

	if err == sql.ErrNoRows {
		return nil, errors.New("client not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting client: %v", err)
	}

	var transactionsList []models.Transaction
	if err := json.Unmarshal(transactions, &transactionsList); err != nil {
		return nil, fmt.Errorf("error unmarshaling transactions: %v", err)
	}

	switch clientType {
	case "personal":
		return &models.PersonalClient{
			BaseClient: models.BaseClient{
				ID:           clientID,
				Name:         name,
				Balance:      balance,
				Transactions: transactionsList,
			},
			CPF: cpf.String,
		}, nil
	case "corporate":
		return &models.CorporateClient{
			BaseClient: models.BaseClient{
				ID:           clientID,
				Name:         name,
				Balance:      balance,
				Transactions: transactionsList,
			},
			CNPJ: cnpj.String,
		}, nil
	default:
		return nil, fmt.Errorf("unknown client type: %s", clientType)
	}
}

func (p *PostgresDB) UpdateClient(client models.Client) error {
	var transactions []byte
	var err error

	switch c := client.(type) {
	case *models.PersonalClient:
		transactions, err = json.Marshal(c.Transactions)
		if err != nil {
			return fmt.Errorf("error marshaling transactions: %v", err)
		}

		query := `
			UPDATE clients
			SET balance = $1, transactions = $2
			WHERE id = $3 AND client_type = 'personal'`

		result, err := p.db.Exec(query, c.Balance, transactions, c.ID)
		if err != nil {
			return fmt.Errorf("error updating personal client: %v", err)
		}

		rows, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("error getting rows affected: %v", err)
		}
		if rows == 0 {
			return errors.New("client not found")
		}

	case *models.CorporateClient:
		transactions, err = json.Marshal(c.Transactions)
		if err != nil {
			return fmt.Errorf("error marshaling transactions: %v", err)
		}

		query := `
			UPDATE clients
			SET balance = $1, transactions = $2
			WHERE id = $3 AND client_type = 'corporate'`

		result, err := p.db.Exec(query, c.Balance, transactions, c.ID)
		if err != nil {
			return fmt.Errorf("error updating corporate client: %v", err)
		}

		rows, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("error getting rows affected: %v", err)
		}
		if rows == 0 {
			return errors.New("client not found")
		}

	default:
		return errors.New("invalid client type")
	}

	return nil
}

func (p *PostgresDB) ListClients() ([]models.Client, error) {
	query := `
		SELECT id, name, balance, client_type, cpf, cnpj, transactions
		FROM clients`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error listing clients: %v", err)
	}
	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var (
			id           string
			name         string
			balance      float64
			clientType   string
			cpf          sql.NullString
			cnpj         sql.NullString
			transactions []byte
		)

		if err := rows.Scan(&id, &name, &balance, &clientType, &cpf, &cnpj, &transactions); err != nil {
			return nil, fmt.Errorf("error scanning client: %v", err)
		}

		var transactionsList []models.Transaction
		if err := json.Unmarshal(transactions, &transactionsList); err != nil {
			return nil, fmt.Errorf("error unmarshaling transactions: %v", err)
		}

		switch clientType {
		case "personal":
			clients = append(clients, &models.PersonalClient{
				BaseClient: models.BaseClient{
					ID:           id,
					Name:         name,
					Balance:      balance,
					Transactions: transactionsList,
				},
				CPF: cpf.String,
			})
		case "corporate":
			clients = append(clients, &models.CorporateClient{
				BaseClient: models.BaseClient{
					ID:           id,
					Name:         name,
					Balance:      balance,
					Transactions: transactionsList,
				},
				CNPJ: cnpj.String,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating clients: %v", err)
	}

	return clients, nil
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}
